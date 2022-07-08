package generators

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	_ "github.com/davecgh/go-spew/spew"
)

//go:embed templates/*
var fs embed.FS

type Payload struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GenPage(action string, args []string) ([]byte, error) {

	sourceFileName := args[0]
	_ = sourceFileName
	_ = action

	jsonFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	payload := Payload{}
	ba, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(ba, &payload)

	// spew.Dump(payload)

	tmpl, err := template.ParseFS(fs, "templates/page.tmpl")
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
