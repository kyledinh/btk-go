package codegen

import (
	"encoding/json"
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"
)

func MakeJsonSchemaFromYaml(filePath string) ([]byte, error) {
	swagger, err := LoadSwagger(filePath)

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
