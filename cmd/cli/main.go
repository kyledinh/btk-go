package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/ghodss/yaml"
	"github.com/kyledinh/btk-go/config"
	"github.com/kyledinh/btk-go/pkg/codex"
	"github.com/kyledinh/btk-go/pkg/gencode"
	"github.com/kyledinh/btk-go/pkg/moxerr"
	"github.com/kyledinh/btk-go/pkg/moxutil"
	goyaml "gopkg.in/yaml.v2"
)

func errCheckLogFatal(err error, me *error) {
	if err != nil {
		log.Fatal(moxerr.NewWrappedError(err.Error(), me))
		// panic(err)
	}
}

func main() {

	// PARSE INPUT
	yamltojson := flag.Bool("yamltojson", false, "Convert yaml to json instead of the default json to yaml.")
	y2j := flag.Bool("y2j", false, "Convert yaml to json instead of the default json to yaml.")
	jsontoyaml := flag.Bool("jsontoyaml", false, "Convert default json to yaml.")
	j2y := flag.Bool("j2y", false, "Convert the default json to yaml.")

	gentest := flag.Bool("gentest", false, "Generate a test file")
	genFlag := flag.String("gen", "", "Generator with an input")

	docsFlag := flag.Bool("docs", false, "Output a documentation file")
	snipFlag := flag.Bool("snip", false, "Output a snip/snippet")
	versionFlag := flag.Bool("v", false, "-v for version")

	yaml2goschema := flag.String("yaml2goschema", "", "Convert spec.yaml/your-spec.yaml to go schema.")

	inputfile := flag.String("i", "", "Specify a spec yaml file  '-i=spec.yaml'")
	outfile := flag.String("o", "", "Specify a file to write to instead of STDOUT,  '-o=filename.ext'")
	destdir := flag.String("d", "", "Specify a directory to write to instead of ./,  '-d=output'")

	flag.Parse()
	args := flag.Args()
	_ = args
	_ = genFlag

	// Don't wrap long lines
	goyaml.FutureLineWrap()

	var (
		outBytes []byte
		err      error
	)

	// MAIN SWITCH
	if *versionFlag {
		outBytes = []byte(`btk-cli ` + config.Version)
		os.Stdout.Write(outBytes)
		os.Exit(0)
	}

	if *yaml2goschema != "" {
		outBytes, err = gencode.MakeJsonSchemaFromYaml(*yaml2goschema)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *yamltojson || *y2j {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		errCheckLogFatal(err, &moxerr.ErrResourceRead)

		raw, err := yaml.YAMLToJSON(inBytes)
		outBytes = moxutil.SimplePrettyJson(raw)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *jsontoyaml || *j2y {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		errCheckLogFatal(err, &moxerr.ErrResourceRead)

		outBytes, err = yaml.JSONToYAML(inBytes)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *gentest {
		outBytes, err = codex.GenPage("genpage", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	if *genFlag != "" && *inputfile != "" {
		var config codegen.Configuration
		config.PackageName = *genFlag
		config.Generate.Models = true

		log.Println("inputfile: ", *inputfile)
		err := gencode.GenerateModels(*inputfile, *destdir, config)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *docsFlag {
		outBytes, err = codex.GetDoc("stdout", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	if *snipFlag {
		outBytes, err = codex.GetSnip("stdout", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	outBytes = moxutil.PrettyJsonBytes(outBytes)

	// OUTPUT to file or STDOUT
	if *outfile != "" {
		err := ioutil.WriteFile(*outfile, outBytes, 0755)
		errCheckLogFatal(err, &moxerr.ErrWriteFile)
	} else {
		os.Stdout.Write(outBytes)
	}

}
