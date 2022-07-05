package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	goyaml "gopkg.in/yaml.v2"
)

func main() {
	yamltojson := flag.Bool("yamltojson", false, "Convert yaml to json instead of the default json to yaml.")
	y2j := flag.Bool("y2j", false, "Convert yaml to json instead of the default json to yaml.")
	jsontoyaml := flag.Bool("jsontoyaml", false, "Convert default json to yaml.")
	j2y := flag.Bool("j2y", false, "Convert the default json to yaml.")
	flag.Parse()

	// Don't wrap long lines
	goyaml.FutureLineWrap()

	inBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var outBytes []byte

	if *yamltojson || *y2j {
		outBytes, err = yaml.YAMLToJSON(inBytes)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(outBytes)
	}

	if *jsontoyaml || *j2y {
		outBytes, err = yaml.JSONToYAML(inBytes)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(outBytes)
	}
	os.Stdout.Write(outBytes)
}
