package prefab

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	_ "github.com/davecgh/go-spew/spew"
)

//go:embed templates/*
var prefabFS embed.FS

type payload struct {
	ba []byte
}

func GetBytesTemplate(action string, args []string) ([]byte, error) {

	_ = action
	payload := payload{}

	targetPrefab := "default"
	if len(args) >= 1 {
		targetPrefab = args[0]
	}

	var templateFilename string

	switch targetPrefab {
	case "unit-test":
		templateFilename = "templates/unit-test.go"
	case "readme":
		templateFilename = "templates/readme.md"
	default:
		templateFilename = "templates/help.md"
	}

	tmpl, err := template.ParseFS(prefabFS, templateFilename)
	if err != nil {
		fmt.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, payload)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes(), err
}
