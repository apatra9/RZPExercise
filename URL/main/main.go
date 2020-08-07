package main

import (
	"flag"

	"fmt"
	"net/http"

	"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	//define flags command line
	yamlfp := flag.String("yamlFile", "yaml", "path of yaml file")
	jsonfp := flag.String("jsonFile", "json", "path of json file")
	flag.Parse()
	yamlfile := ioUtil.ReadFile(*yamlfp)
	jsonfile := ioUtil.ReadFile(*jsonfp)
	var err error

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	Handler := mapHandler //default

	// Build the YAMLHandler and JSONHandler using the mapHandler as the
	// fallback

	if jsonfile != nil {
		Handler, err = urlshort.JSONHandler(jsonfile, mapHandler)
	} else if yamlfile != nil {
		Handler, err = urlshort.YAMLHandler(yamlfile, mapHandler)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", Handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
