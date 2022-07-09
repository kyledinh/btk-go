package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	generators "github.com/kyledinh/btk-go/internal/generator"
	"github.com/kyledinh/btk-go/pkg/prefab"
	goyaml "gopkg.in/yaml.v2"
)

func main() {
	yamltojson := flag.Bool("yamltojson", false, "Convert yaml to json instead of the default json to yaml.")
	y2j := flag.Bool("y2j", false, "Convert yaml to json instead of the default json to yaml.")
	jsontoyaml := flag.Bool("jsontoyaml", false, "Convert default json to yaml.")
	j2y := flag.Bool("j2y", false, "Convert the default json to yaml.")

	gentest := flag.Bool("gentest", false, "Generate a test file")
	prefabFlag := flag.Bool("prefab", false, "Output a prefab file")
	outfile := flag.String("outfile", "", "")

	flag.Parse()
	args := flag.Args()
	_ = args

	// Don't wrap long lines
	goyaml.FutureLineWrap()

	var (
		outBytes []byte
		err      error
	)

	if *yamltojson || *y2j {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		outBytes, err = yaml.YAMLToJSON(inBytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *jsontoyaml || *j2y {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		outBytes, err = yaml.JSONToYAML(inBytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *gentest {
		outBytes, err = generators.GenPage("genpage", args)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *prefabFlag {
		outBytes, err = prefab.GetBytesTemplate("stdout", args)
		if err != nil {
			log.Fatal(err)
		}
	}

	// OUTPUT Stdout or file
	if *outfile != "" {

	} else {
		os.Stdout.Write(outBytes)
	}

}
