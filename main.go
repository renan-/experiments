package main

import (
	_ "net/http"

	_ "github.com/codegangsta/negroni"
	_ "github.com/gorilla/mux"
	_ "github.com/jteeuwen/go-bindata"
)

//
// @since always
//
func main() {
	NewServer().Run(":80")
}
