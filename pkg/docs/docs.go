package docs

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	_ "github.com/davecgh/go-spew/spew"
)

//go:embed files/*
var docsFS embed.FS

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
	case "coding-guide":
		templateFilename = "files/coding-guide.md"
	case "resources":
		templateFilename = "files/resources.md"
	default:
		templateFilename = "files/help.md"
	}

	tmpl, err := template.ParseFS(docsFS, templateFilename)
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
