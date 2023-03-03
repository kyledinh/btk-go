package server

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/kyledinh/btk-go/pkg/codex"
)

func docsHandler(w http.ResponseWriter, r *http.Request) {

	if strings.HasSuffix(string(r.URL.Path), ".md") {
		lastSlash := strings.LastIndex((r.URL.Path), "/")
		filename := string(r.URL.Path)[lastSlash+1:]

		args := []string{filename}
		md, err := codex.GetDoc("stdout", args)
		if err != nil {
			panic(err)
		}

		opts := html.RendererOptions{
			Flags:          html.FlagsNone,
			RenderNodeHook: renderHook,
		}
		renderer := html.NewRenderer(opts)
		output := string(markdown.ToHTML(md, nil, renderer))

		t := template.New("Render")
		t, err = t.Parse(htmlHeader + "{{.}}" + htmlFooter)
		t = template.Must(t, err)

		var processed bytes.Buffer
		t.Execute(&processed, output)

		w.Write(processed.Bytes())

	} else {
		http.FileServer(http.FS(codex.DocsFS)).ServeHTTP(w, r)
	}
}

var htmlHeader = `
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
   <title>Blog Post Example</title>
</head>
<body>
<br>
<div class="container is-max-desktop">
<div class="content">`

var htmlFooter = `
</div>
</div>
<br>
</body>
</html>`

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {

	if _, ok := node.(*ast.Heading); ok {
		level := strconv.Itoa(node.(*ast.Heading).Level)

		if entering && level == "1" {
			w.Write([]byte(`<h1 class="title is-1 has-text-centered">`))
		} else if entering {
			w.Write([]byte("<h" + level + ">"))
		} else {
			w.Write([]byte("</h" + level + ">"))
		}

		return ast.GoToNext, true

	} else if _, ok := node.(*ast.Image); ok {
		src := string(node.(*ast.Image).Destination)

		c := node.(*ast.Image).GetChildren()[0]
		alt := string(c.AsLeaf().Literal)

		if entering && alt != "" {
			w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `" alt="` + alt + `">`))
		} else if entering {
			w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `">`))
		} else {
			w.Write([]byte(`</figure>`))
		}

		return ast.SkipChildren, true
	} else {
		return ast.GoToNext, false
	}
}
