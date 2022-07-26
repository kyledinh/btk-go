package gencode

import (
	"os"

	"github.com/cweill/gotests/gotests/process"
)

func GenerateTests(srcFile string, destDir string, opts Config) error {
	var err error
	srcFiles := make([]string, 0)

	if destDir == "" {
		destDir = "."
	}

	// WRAP the cweill/gotests process.Run
	// TODO convert opts to process.Options
	opt := process.Options{
		Parallel:    true,
		WriteOutput: true,
		AllFuncs:    true,
	}
	srcFiles = append(srcFiles, srcFile)

	process.Run(os.Stdout, srcFiles, &opt)

	return err
}
