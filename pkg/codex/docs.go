package codex

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"

	_ "github.com/davecgh/go-spew/spew"
)

//go:embed docs/*

var docsFS embed.FS

func GetDoc(action string, args []string) ([]byte, error) {

	_ = action
	payload := payload{}

	targetPrefab := "default"
	if len(args) >= 1 {
		targetPrefab = args[0]
	}

	var templateFilename string
	var availabeFiles []string

	files, err := docsFS.ReadDir("docs")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Name() == targetPrefab {
			templateFilename = "docs/" + file.Name()
		}
		if !file.IsDir() {
			availabeFiles = append(availabeFiles, file.Name())
		}
	}

	// NO MATCH FOR TEMPLATE, SEND A LIST OF AVAILABLE FILES
	if templateFilename == "" {
		templateFilename = strings.Join(availabeFiles, "\n")
		return []byte(templateFilename), err
	}

	// PARSE THE TARGET FILE
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
