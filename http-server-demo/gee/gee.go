package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(rw, "404 NOT FOUND: %s\n", r.URL)

	}
}

func main() {
	engine := new(Engine)

	http.ListenAndServe(":8081", engine)
}
