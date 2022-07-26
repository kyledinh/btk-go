package gencode

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/url"
	"strings"
	"text/template"
	"unicode"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
)

// Embed the templates directory
//go:embed templates
var TemplatesFS embed.FS

type Payload struct {
	SchemaName     string
	ModuleName     string
	Package        string
	GenVersion     string
	SpecVersion    string
	SpecFile       string
	GoSchema       codegen.Schema
	Imports        []string
	PubFieldLookup map[string]string
	PubFieldName   func(string) string
	FilterGoType   func(string) string
	FilterGoImport func(string) string
}

type Config struct {
	Options string // TODO: fill in fields as needed
}

func MakeJsonSchemaFromYaml(filePath string) ([]byte, error) {
	swagger, err := LoadSwagger(filePath)
	if err != nil {
		return nil, err
	}

	outBytes, err := json.Marshal(swagger)
	return outBytes, err
}

func LoadSwagger(filePath string) (swagger *openapi3.T, err error) {

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	u, err := url.Parse(filePath)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return loader.LoadFromURI(u)
	} else {
		return loader.LoadFromFile(filePath)
	}
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

func PascalFrom_snake_case(snake_case string) string {
	var PascalStr string = ""
	snakes := strings.Split(snake_case, "_")
	for _, snake := range snakes {
		PascalStr += PubCapitalize(snake)
	}
	return PascalStr
}

func FilterGoType(str string) string {
	if str == "openapi_types.UUID" {
		str = "uuid.UUID"
	}
	return str
}

func FilterGoImport(str string) string {
	if strings.Contains(str, "UUID") {
		return "github.com/google/uuid"
	}
	if str == "time.Time" {
		return "time"
	}
	return ""
}
