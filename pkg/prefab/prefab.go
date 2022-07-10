package prefab

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"

	_ "github.com/davecgh/go-spew/spew"
)

//go:embed templates/*
var prefabFS embed.FS

type payload struct {
	ba []byte
}

func KeywordFromFilename(filename string) string {
	keyword := filename[strings.Index(filename, "/")+1:]
	return keyword
}

func GetBytesTemplate(action string, args []string) ([]byte, error) {

	_ = action
	payload := payload{}

	targetPrefab := "default"
	if len(args) >= 1 {
		targetPrefab = args[0]
	}

	var templateFilename string
	var availabeFiles []string

	files, err := prefabFS.ReadDir("templates")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Name() == targetPrefab {
			templateFilename = "templates/" + file.Name()
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

	// PARSE THE TARGE FILE
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
