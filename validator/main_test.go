package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kaptinlin/jsonschema"
)

func TestValidateGoodFiles(t *testing.T) {
	// Use real schema and real data files
	baseDir := ".."

	tests := []struct {
		name       string
		schemaFile string
		dataFile   string
	}{
		{
			name:       "burdah fasl v1.0",
			schemaFile: "burdah_fasl_v1.0.schema.json",
			dataFile:   "sholawat/burdah/nu_online/fasl/1.json",
		},
		{
			name:       "burdah fasl v1.1",
			schemaFile: "burdah_fasl_v1.1.schema.json",
			dataFile:   "sholawat/burdah/minhaj_ul_quran/fasl/1.json",
		},
		{
			name:       "tunggal v1.1",
			schemaFile: "sholawat_tunggal_v1.1.schema.json",
			dataFile:   "sholawat/tunggal/addinu-lana.json",
		},
		{
			name:       "suluk v1.1",
			schemaFile: "suluk_v1.1.schema.json",
			dataFile:   "sholawat/suluk/adimishsholata-alal-habib.json",
		},
		{
			name:       "diba fasl v1.0",
			schemaFile: "diba_fasl_v1.0.schema.json",
			dataFile:   "sholawat/diba/nu_online/fasl/1.json",
		},
		{
			name:       "simtudduror fasl v1.0",
			schemaFile: "simtudduror_fasl_v1.0.schema.json",
			dataFile:   "sholawat/simtudduror/nu_online/fasl/1.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaPath := filepath.Join(baseDir, "schemas", tt.schemaFile)
			schemaData, err := os.ReadFile(schemaPath)
			if err != nil {
				t.Fatalf("failed to read schema %s: %v", schemaPath, err)
			}

			compiler := jsonschema.NewCompiler()
			schema, err := compiler.Compile(schemaData)
			if err != nil {
				t.Fatalf("failed to compile schema %s: %v", schemaPath, err)
			}

			dataPath := filepath.Join(baseDir, tt.dataFile)
			jsonData, err := os.ReadFile(dataPath)
			if err != nil {
				t.Fatalf("failed to read data file %s: %v", dataPath, err)
			}

			var instance map[string]any
			if err := json.Unmarshal(jsonData, &instance); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			result := schema.Validate(instance)
			if !result.IsValid() {
				var msgs []string
				for _, detail := range result.ToList().Details {
					if !detail.Valid {
						for _, e := range detail.Errors {
							msgs = append(msgs, e)
						}
					}
				}
				t.Errorf("validation failed: %s", strings.Join(msgs, "; "))
			}
		})
	}
}

func TestValidateBadJSON(t *testing.T) {
	baseDir := ".."
	schemaPath := filepath.Join(baseDir, "schemas", "sholawat_tunggal_v1.1.schema.json")
	schemaData, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("failed to read schema: %v", err)
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(schemaData)
	if err != nil {
		t.Fatalf("failed to compile schema: %v", err)
	}

	t.Run("missing required name field", func(t *testing.T) {
		// Missing required "name" field
		instance := map[string]any{
			"text": map[string]any{
				"1": map[string]any{"arabic": "بسم الله"},
			},
			"last_updated": "2024-01-01",
		}
		result := schema.Validate(instance)
		if result.IsValid() {
			t.Error("expected validation failure for missing 'name' field")
		}
	})

	t.Run("invalid last_updated format", func(t *testing.T) {
		instance := map[string]any{
			"name":         "test",
			"text":         map[string]any{"1": map[string]any{"arabic": "بسم الله"}},
			"last_updated": "01/01/2024",
		}
		result := schema.Validate(instance)
		if result.IsValid() {
			t.Error("expected validation failure for invalid date format")
		}
	})

	t.Run("non-numeric text key", func(t *testing.T) {
		instance := map[string]any{
			"name": "test",
			"text": map[string]any{
				"abc": map[string]any{"arabic": "بسم الله"},
			},
			"last_updated": "2024-01-01",
		}
		result := schema.Validate(instance)
		if result.IsValid() {
			t.Error("expected validation failure for non-numeric text key")
		}
	})
}
