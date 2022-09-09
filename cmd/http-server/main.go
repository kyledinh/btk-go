package main

import (
	"fmt"
	"net/http"

	"github.com/kyledinh/btk-go/pkg/codex"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.Handle("/docs/", http.FileServer(http.FS(codex.DocsFS)))
	http.Handle("/snippets/", http.FileServer(http.FS(codex.SnippetsFS)))
	http.Handle("/templates/", http.FileServer(http.FS(codex.TemplatesFS)))

	fmt.Printf("Starting HTTP server on port %s.\n", "8001")
	http.ListenAndServe(":8001", nil)

	select {}
}
