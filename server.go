package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/zemirco/couchdb"
)

// A wrapper
type Server struct {
	// Router
	router *mux.Router

	// Negroni
	ng *negroni.Negroni
}

var (
	CouchDB *couchdb.Client = nil
	err     error
)

func NewServer() *Server {
	srv := Server{}
	// The router for this Server
	srv.router = mux.NewRouter()

	// Negroni, courtesy of codegangsta
	srv.ng = negroni.Classic()

	// API takes care of all resources
	// It is located on a separate subdomain
	api := srv.router.Host("api.localhost").Subrouter()
	api.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")
	api.HandleFunc("/{resource}", CollectionsHandler)
	api.HandleFunc("/{resource}/{id}", ModelsHandler)

	// Serve static files from memory
	fs := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir})
	srv.router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// Use GZIP compression middleware
	srv.ng.Use(gzip.Gzip(gzip.DefaultCompression))
	// Setup negroni to use our routes
	srv.ng.UseHandler(srv.router)

	if CouchDB == nil {
		CouchDB, err = couchdb.NewClient("http://127.0.0.1:5984/")
		if err != nil {
			panic(err)
		}
	}

	return &srv
}

func (s *Server) Run(addr string) {
	// Run negroni
	s.ng.Run(addr)
}
