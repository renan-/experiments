package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	//"github.com/boltdb/bolt"
	"github.com/codegangsta/negroni"
	"github.com/couchbaselabs/gocb"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
)

var (
	err error
	Cluster *gocb.Cluster
	once sync.Once
)

// A wrapper
type Server struct {
	// Router
	router *mux.Router

	// Negroni
	ng *negroni.Negroni
}

func NewServer() *Server {
	// First of all initialize database
	once.Do(func() {
		Cluster, err = gocb.Connect("http://robotox:8091")
		if err != nil {
			log.Fatal(err)
		}

		bucket, err := Cluster.OpenBucket("couchbase://localhost/beer-sample", "")
		if err != nil {
			log.Fatal(err)
		}

		var out struct {
			name string `json:"name"`
			city string `json:"city"`
		}
		_, err = bucket.Get("21st_amendment_brewery_cafe", out)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v", out)
	})

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
	api.HandleFunc("/{resource}/{id:[0-9]+}", ModelsHandler)

	// Serve static files from memory
	fs := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir})
	srv.router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// Use GZIP compression middleware
	srv.ng.Use(gzip.Gzip(gzip.DefaultCompression))
	// Setup negroni to use our routes
	srv.ng.UseHandler(srv.router)

	return &srv
}

func (s *Server) Run(addr string) {
	// Run negroni
	s.ng.Run(addr)
}
