package gencode

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
	"text/template"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kyledinh/btk-go/config"
	"github.com/kyledinh/btk-go/ignored/codegen"
)

// Embed the templates directory
//go:embed templates
var TemplatesFS embed.FS

type Payload struct {
	SchemaName     string
	ModuleName     string
	Version        string
	GoSchema       codegen.Schema
	PubFieldLookup map[string]string
	PubFieldName   func(string) string
	FilterGoType   func(string) string
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

		var pubFieldLookup = make(map[string]string, len(goSchema.Properties))

		for _, field := range goSchema.Properties {
			pubFieldLookup[field.JsonFieldName] = strings.ToUpper(field.JsonFieldName)
		}

		var payload = Payload{
			SchemaName:     schemaName,
			ModuleName:     "github.com/kyledinh/btk-cli-go",
			Version:        config.Version,
			GoSchema:       goSchema,
			PubFieldLookup: pubFieldLookup,
			PubFieldName:   PubCapitalize,
			FilterGoType:   FilterGoType,
		}

		tmpl, err := template.ParseFS(TemplatesFS, "templates/model.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, payload)
		if err != nil {
			fmt.Println(err)
		}
		outBytes := buf.Bytes()

		err = ioutil.WriteFile("dist/"+filename, outBytes, 0755)
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

func PubCapitalize(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func FilterGoType(str string) string {
	if str == "openapi_types.UUID" {
		str = "uuid.UUID"
	}

	return str
}
