package main

import "github.com/zemirco/couchdb"

// A single user
type User struct {
	couchdb.Document

	Firstname string `json:"firstname"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
