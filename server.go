package main

import (
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"net/http"
)

func main() {
	r := render.New(render.Options{
		Directory: "templates",
		Extensions: []string{".html"},
		IsDevelopment: true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3001")
}
