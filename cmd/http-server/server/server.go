package server

import (
	"fmt"
	"net/http"

	"github.com/kyledinh/btk-go/pkg/codex"
	"github.com/kyledinh/btk-go/config"
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

func Server() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/docs/", docsHandler)
	mux.Handle("/snippets/", http.FileServer(http.FS(codex.SnippetsFS)))
	mux.Handle("/templates/", http.FileServer(http.FS(codex.TemplatesFS)))

	fmt.Printf("Starting HTTP server on port %s.\n", "8001")
	fmt.Printf("Running version: %s\n\n",  config.Version)
	fmt.Println("Endpoints:")
	fmt.Println("http://localhost:8001/docs/")
	fmt.Println("http://localhost:8001/snippets/")
	fmt.Println("http://localhost:8001/templates/")
	fmt.Println("http://localhost:8001/headers")
	http.ListenAndServe(":8001", mux)

	select {}
}
