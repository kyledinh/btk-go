package generators

import (
	"embed"
	"encoding/json"
	"flag"
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

func main() {
	action := flag.Bool("action", false, "trigger action")
	flag.Parse()
	args := flag.Args()

	sourceFileName := args[0]
	_ = sourceFileName

	if *action {
		fmt.Println("Action triggered")
	}

	jsonFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	payload := Payload{}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(bytes, &payload)

	// spew.Dump(payload)

	tmpl, err := template.ParseFS(fs, "templates/page.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(os.Stdout, payload)
	if err != nil {
		fmt.Println(err)
	}
}
