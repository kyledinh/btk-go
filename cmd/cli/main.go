package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/ghodss/yaml"
	"github.com/kyledinh/btk-go/config"
	webserver "github.com/kyledinh/btk-go/cmd/http-server/server"
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

	genTest := flag.Bool("gentest", false, "Generate a unit test scaffolding '-gentest -i file.go' ")
	genModels := flag.String("gen", "", "Generate models '-gen=model i=specs/project.yaml -d=internal/model' ")

	docsFlag := flag.Bool("docs", false, "Output a documentation file")
	snipFlag := flag.Bool("snip", false, "Output a snip/snippet")
	versionFlag := flag.Bool("v", false, "-v for version")

	webFlag := flag.Bool("web", false, "-web to launch http server")


	yaml2goschema := flag.String("yaml2goschema", "", "Convert spec.yaml/your-spec.yaml to go schema.")

	inputfile := flag.String("i", "", "Specify a spec yaml file  '-i=spec.yaml'")
	outfile := flag.String("o", "", "Specify a file to write to instead of STDOUT,  '-o=filename.ext'")
	destdir := flag.String("d", "", "Specify a directory to write to instead of ./,  '-d=output'")

	flag.Parse()
	args := flag.Args()
	_ = args
	_ = genModels

	// Don't wrap long lines
	goyaml.FutureLineWrap()

	var (
		outBytes []byte
		err      error
	)

	// MAIN SWITCH
	if *webFlag {
		webserver.Server()
		select {}
	}

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

		outBytes, err = yaml.JSONToYAML(inBytes)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *genTest && *inputfile != "" {
		var cfg gencode.Config
		// outBytes, err = codex.GenPage("genpage", args)
		err := gencode.GenerateTests(*inputfile, *destdir, cfg)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	if *genModels != "" && *inputfile != "" {
		var config codegen.Configuration
		config.PackageName = *genModels
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
