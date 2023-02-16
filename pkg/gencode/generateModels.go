package gencode

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kyledinh/btk-go/config"
)

func GenerateModels(specFile string, destDir string, opts codegen.Configuration) error {

	if destDir == "" {
		destDir = "."
	}

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

		imports := make([]string, 0)
		for _, field := range goSchema.Properties {
			pubFieldLookup[field.JsonFieldName] = strings.ToUpper(field.JsonFieldName)
			entry := FilterGoImport(field.GoTypeDef())
			if entry != "" {
				imports = append(imports, entry)
			}
		}

		var payload = Payload{
			SchemaName:     PascalFrom_snake_case(schemaName),
			ModuleName:     "github.com/kyledinh/btk-cli-go",
			GenVersion:     config.Version,
			SpecVersion:    spec.Info.Version,
			SpecFile:       specFile,
			Package:        "model",
			GoSchema:       goSchema,
			Imports:        imports,
			PubFieldLookup: pubFieldLookup,
			PubFieldName:   PascalFrom_snake_case,
			FilterGoType:   FilterGoType,
			FilterGoImport: FilterGoImport,
		}

		// Generate Golang models
		tmpl, err := template.ParseFS(TemplatesFS, "templates/model.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, payload)
		if err != nil {
			fmt.Println(err)
		}

		outBytes, err := format.Source(buf.Bytes())
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(destDir+"/"+filename, outBytes, 0755)
		if err != nil {
			return err
		}

		// ALSO Generate Scala 3 models!!!
		filename = "gen.model." + strings.ToLower(schemaName) + ".scala"

		tmplScala, err := template.ParseFS(TemplatesFS, "templates/scala-3-model.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		var bufScala bytes.Buffer
		err = tmplScala.Execute(&bufScala, payload)
		if err != nil {
			fmt.Println(err)
		}

		outBytesScala, err := format.Source(bufScala.Bytes())
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(destDir+"/"+filename, outBytesScala, 0755)
		if err != nil {
			return err
		}

	}

	return err
}

func Generate(spec *openapi3.T, opts codegen.Configuration) (string, error) {
	return "", nil
}
