package codex

import "strings"

type payload struct {
	ba []byte
}

func KeywordFromFilename(filename string) string {
	keyword := filename[strings.Index(filename, "/")+1:]
	return keyword
}
