// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Unique id of the pet
	Id openapi_types.UUID `json:"id"`

	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody = NewPet

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = AddPetJSONBody

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"+1sAAAD//wYFIpNREgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
