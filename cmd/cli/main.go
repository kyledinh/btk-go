package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	generators "github.com/kyledinh/btk-go/internal/generator"
	"github.com/kyledinh/btk-go/pkg/docs"
	"github.com/kyledinh/btk-go/pkg/moxerr"
	"github.com/kyledinh/btk-go/pkg/prefab"
	goyaml "gopkg.in/yaml.v2"
)

func errCheckLogFatal(err error, me *error) {
	if err != nil {
		log.Fatal(moxerr.NewWrappedError(err.Error(), me))
		// panic(err)
	}
}

func main() {
	yamltojson := flag.Bool("yamltojson", false, "Convert yaml to json instead of the default json to yaml.")
	y2j := flag.Bool("y2j", false, "Convert yaml to json instead of the default json to yaml.")
	jsontoyaml := flag.Bool("jsontoyaml", false, "Convert default json to yaml.")
	j2y := flag.Bool("j2y", false, "Convert the default json to yaml.")

	gentest := flag.Bool("gentest", false, "Generate a test file")
	docsFlag := flag.Bool("docs", false, "Output a documentation file")
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
		errCheckLogFatal(err, &moxerr.ErrResourceRead)

		outBytes, err = yaml.YAMLToJSON(inBytes)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *jsontoyaml || *j2y {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		errCheckLogFatal(err, &moxerr.ErrResourceRead)

		outBytes, err = yaml.JSONToYAML(inBytes)
		errCheckLogFatal(err, &moxerr.ErrConversionFormat)
	}

	if *gentest {
		outBytes, err = generators.GenPage("genpage", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	if *docsFlag {
		outBytes, err = docs.GetBytesTemplate("stdout", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	if *prefabFlag {
		outBytes, err = prefab.GetBytesTemplate("stdout", args)
		errCheckLogFatal(err, &moxerr.ErrCLIAction)
	}

	// OUTPUT Stdout or file
	if *outfile != "" {

	} else {
		os.Stdout.Write(outBytes)
	}

}
