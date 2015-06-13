package main

import (
	"encoding/json"
	"fmt"
	//"log"
    //"mime"
	"net/http"
	//"path/filepath"

	//"github.com/boltdb/bolt"
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"github.com/jteeuwen/go-bindata"
)

// Simple enough
func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, string(MustAsset("index.html")))
}

//
// Handles API models GET, POST, UPDATE, DELETE
//
func ModelsHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resource_name := vars["resource"]
	//resource_id   := vars["id"]

	var resource Resource

	switch resource_name {
	case "users":
		resource = NewUser("toto", "toto@localhost", "toto42")
	default:
		resource = struct{}{}
	}

	contents, err := json.MarshalIndent(resource, "", "    ")
	if err != nil {
		panic(err)
	}

	rw.Header().Add("Content-type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "http://localhost")

	fmt.Fprintf(rw, string(contents))
}

//
// Handle API collections GET POST
//
func CollectionsHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resource_name := vars["resource"]

	fmt.Fprintf(rw, resource_name)
}

/*
// This function returns a http.HandlerFunc
// that handles the asset provided by provider.
// It adds Content-Type header based on file extension
//
func AssetHandler(provider func() (*asset, error)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		contents, err := provider()
		if err != nil {
			panic(err)
		}

		Dynamically determine content type
		var content_type string

		switch path.Ext(r.URL.Path) {
		case ".css":
			content_type = "text/css"
		case ".js":
			content_type = "text/javascript"
		case ".html":
		default:
			content_type = "text/html"
		}

		rw.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(r.URL.Path)))

		fmt.Fprintf(rw, string(contents.bytes))
	}
}
*/