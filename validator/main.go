package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

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

type cachedSchema struct {
	schema *jsonschema.Schema
	mu     sync.Mutex
}

func main() {
	sources, err := loadSources("./../sholawat/sholawat.json")
	if err != nil {
		log.Fatalf("Error loading sources: %v", err)
	}

	baseDir := "./../"
	var validationFailed int32

	registeredFiles := collectRegisteredFiles(sources)
	reportOrphanFiles(baseDir, registeredFiles)
	schemaCache := compileSchemaCache(sources, baseDir, &validationFailed)
	validateAllSources(sources, schemaCache, baseDir, &validationFailed)

	if atomic.LoadInt32(&validationFailed) > 0 {
		os.Exit(1)
	}
}

func loadSources(registryPath string) ([]Sholawat, error) {
	data, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, err
	}

	var sources []Sholawat
	if err := json.Unmarshal(data, &sources); err != nil {
		return nil, err
	}

	return sources, nil
}

func collectRegisteredFiles(sources []Sholawat) map[string]bool {
	registeredFiles := make(map[string]bool)
	for _, sholawat := range sources {
		for _, source := range sholawat.Sources {
			for _, file := range source.Files {
				fullPath := filepath.Join(source.PathFiles, file)
				registeredFiles[fullPath] = true
			}
		}
	}

	return registeredFiles
}

func reportOrphanFiles(baseDir string, registeredFiles map[string]bool) {
	orphanFiles := findOrphanFiles(baseDir, registeredFiles)
	if len(orphanFiles) == 0 {
		return
	}

	log.Printf("\n=== ORPHAN FILES (not in registry) ===")
	for _, file := range orphanFiles {
		log.Printf("Orphan file: %s\n", file)
	}
	log.Printf("Total orphan files: %d\n", len(orphanFiles))
}

func compileSchemaCache(sources []Sholawat, baseDir string, validationFailed *int32) map[string]*cachedSchema {
	schemaCache := make(map[string]*cachedSchema)
	for _, sholawat := range sources {
		for _, source := range sholawat.Sources {
			if _, ok := schemaCache[source.Schema]; ok {
				continue
			}

			schemaPath := filepath.Join(baseDir, "schemas", source.Schema)
			schemaData, err := os.ReadFile(schemaPath)
			if err != nil {
				log.Printf("Error reading schema %s: %v", source.Schema, err)
				atomic.AddInt32(validationFailed, 1)
				continue
			}

			compiler := jsonschema.NewCompiler()
			compiled, err := compiler.Compile(schemaData)
			if err != nil {
				log.Printf("Error compiling schema %s: %v", source.Schema, err)
				atomic.AddInt32(validationFailed, 1)
				continue
			}

			schemaCache[source.Schema] = &cachedSchema{schema: compiled}
		}
	}

	return schemaCache
}

func validateAllSources(sources []Sholawat, schemaCache map[string]*cachedSchema, baseDir string, validationFailed *int32) {
	var wg sync.WaitGroup

	for _, sholawat := range sources {
		for _, source := range sholawat.Sources {
			schema, ok := schemaCache[source.Schema]
			if !ok {
				log.Printf("Skipping source %s: schema %s failed to compile", source.SourceName, source.Schema)
				continue
			}

			wg.Add(1)
			go func(source Source, schema *cachedSchema) {
				defer wg.Done()
				validateSourceFile(schema, baseDir, source, validationFailed)
			}(source, schema)
		}
	}

	wg.Wait()
}

func findOrphanFiles(baseDir string, registeredFiles map[string]bool) []string {
	var orphans []string
	sholawatDir := filepath.Join(baseDir, "sholawat")

	err := filepath.Walk(sholawatDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only check JSON files, skip the registry file itself
		if info.IsDir() || !strings.HasSuffix(path, ".json") || path == filepath.Join(sholawatDir, "sholawat.json") {
			return nil
		}

		// Get relative path from sholawat directory
		relPath, err := filepath.Rel(sholawatDir, path)
		if err != nil {
			return nil
		}

		// Normalize path separators
		relPath = filepath.ToSlash(relPath)

		// Also check with "sholawat/" prefix (registry format)
		regPath := "sholawat/" + relPath

		if !registeredFiles[relPath] && !registeredFiles[regPath] {
			orphans = append(orphans, relPath)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking sholawat directory: %v", err)
	}

	return orphans
}

func validateSourceFile(schema *cachedSchema, baseDir string, source Source, validationFailed *int32) {
	for _, file := range source.Files {
		filePath := filepath.Join(baseDir, source.PathFiles, file)
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			atomic.AddInt32(validationFailed, 1)
			continue
		}

		var instance map[string]any
		err = json.Unmarshal(jsonData, &instance)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			atomic.AddInt32(validationFailed, 1)
			continue
		}

		schema.mu.Lock()
		result := schema.schema.Validate(instance)
		schema.mu.Unlock()
		if !result.IsValid() {
			atomic.AddInt32(validationFailed, 1)
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
