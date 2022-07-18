package gencode

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kyledinh/btk-go/ignored/codegen"
)

// Embed the templates directory
//go:embed templates
var templates embed.FS

// globalState stores all global state. Please don't put global state anywhere
// else so that we can easily track it.
var globalState struct {
	options codegen.Configuration
	spec    *openapi3.T
}

func GenerateModels(specFile string, opts codegen.Configuration) error {
	spec, err := LoadSwagger(specFile)
	if err != nil {
		return err
	}

	// Get the Schema keys in Schema, then iterate and fetch the schema
	schemas := spec.Components.Schemas
	for _, schemaName := range codegen.SortedSchemaKeys(schemas) {
		schemaRef := schemas[schemaName]

		goSchema, err := codegen.GenerateGoSchema(schemaRef, []string{schemaName})
		_ = goSchema
		if err != nil {
			return fmt.Errorf("error converting Schema %s to Go type: %w", schemaName, err)
		}
		filename := "gen.model." + strings.ToLower(schemaName) + ".go"
		outBytes := []byte(
			`package model
		
		`)

		//TODO: use template to generate model files

		err = ioutil.WriteFile(filename, outBytes, 0755)
		if err != nil {
			return err
		}

	}

	return err
}

func Generate(spec *openapi3.T, opts codegen.Configuration) (string, error) {
	return "", nil
}

func LoadTemplates(src embed.FS, t *template.Template) error {
	return fs.WalkDir(src, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		buf, err := src.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", path, err)
		}

		templateName := strings.TrimPrefix(path, "templates/")
		tmpl := t.New(templateName)
		_, err = tmpl.Parse(string(buf))
		if err != nil {
			return fmt.Errorf("parsing template '%s': %w", path, err)
		}
		return nil
	})
}

func SanitizeCode(goCode string) string {
	// remove any byte-order-marks which break Go-Code
	// See: https://groups.google.com/forum/#!topic/golang-nuts/OToNIPdfkks
	return strings.Replace(goCode, "\uFEFF", "", -1)
}
