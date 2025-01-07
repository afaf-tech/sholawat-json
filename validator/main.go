package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/kaptinlin/jsonschema"
)

type Source struct {
	SourceName  string   `json:"source_name"`
	Description string   `json:"description"`
	PathFiles   string   `json:"path_files"`
	Files       []string `json:"files"`
	Schema      string   `json:"schema"`
}

type Sholawat struct {
	Name    string   `json:"name"`
	Sources []Source `json:"sources"`
}

type Sources struct {
	Sholawats []Sholawat `json:"sholawat"`
}

func main() {
	data, err := os.ReadFile("./../sholawat/sholawat.json")
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	var sources []Sholawat
	err = json.Unmarshal(data, &sources)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	baseDir := "./../"

	var wg sync.WaitGroup

	for _, sholawat := range sources {
		for _, source := range sholawat.Sources {
			wg.Add(1)

			go func(source Source) {
				defer wg.Done()

				schemaPath := filepath.Join(baseDir, "schemas", source.Schema)
				schemaData, err := os.ReadFile(schemaPath)
				if err != nil {
					log.Printf("Error reading schema for source %s: %v", source.SourceName, err)
					return
				}

				compiler := jsonschema.NewCompiler()
				schema, err := compiler.Compile(schemaData)
				if err != nil {
					log.Printf("Error compiling schema for  %s-%s: %v", sholawat.Name, source.SourceName, err)
					return
				}

				validateSourceFile(schema, baseDir, source)
			}(source)
		}
	}

	wg.Wait()
}

func validateSourceFile(schema *jsonschema.Schema, baseDir string, source Source) {
	for _, file := range source.Files {
		filePath := filepath.Join(baseDir, source.PathFiles, file)
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			continue
		}

		var instance map[string]interface{}
		err = json.Unmarshal(jsonData, &instance)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			return
		}

		result := schema.Validate(instance)
		if !result.IsValid() {
			log.Printf("Validation failed for: %s, file: %s \n", source.SourceName, filePath)
			for _, detail := range result.ToList().Details {
				if !detail.Valid {
					log.Printf("\tEvaluation Path: %s\n", detail.EvaluationPath)

					for _, err := range detail.Errors {
						log.Printf("\t\t- %s\n", err)
					}
				}
			}
		} else {
			log.Printf("Validation passed for file %s\n", filePath)
		}

	}
}
