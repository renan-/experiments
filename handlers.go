package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zemirco/couchdb"
)

//
// Handles API models GET, POST, UPDATE, DELETE
//
func ModelsHandler(rw http.ResponseWriter, r *http.Request) {
	// First retrieve request vars
	vars := mux.Vars(r)
	resource_name := vars["resource"]
	resource_id := vars["id"]

	// Then instantiate the resource accordingly
	var resource couchdb.CouchDoc
	switch resource_name {
	case "users":
		resource = new(User)
	default:
		resource = nil
	}

	// The resource doesn't exist
	if resource == nil {
		return
	}

	// Retrieve the given resource
	database := CouchDB.Use(resource_name)
	database.Get(resource, resource_id)

	// Convert it to JSON
	contents, err := json.MarshalIndent(resource, "", "    ")
	if err != nil {
		panic(err)
	}

	// Set the appropriate headers
	rw.Header().Add("Content-type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "http://localhost")

	// Write JSON
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
