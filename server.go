package main

import (
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var config struct {
	Port int
	Development bool
}

func main() {
	file, e := os.Open("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(file)
  if err := jsonParser.Decode(&config); err != nil {
      fmt.Printf("parsing config file", err.Error())
  }

	r := render.New(render.Options{
		Directory: "templates",
		Extensions: []string{".html"},
		IsDevelopment: config.Development,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + strconv.Itoa(config.Port))
}
