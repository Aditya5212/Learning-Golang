package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is Modules in Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("This is simple Greeting!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is Home Page</h1>"))
}

// go get	           ->    Updates module dependencies and installs packages.
// go install	       ->    Builds and installs packages, especially executables.
// go list -m	       ->    Lists modules and their properties.
// go mod download	   ->    Downloads modules to the local cache.
// go mod edit	       ->    Edits the go.mod file via command line.
// go mod graph	       ->    Prints the module dependency graph.
// go mod init	       ->    Initializes a new go.mod file.
// go mod tidy	       ->    Synchronizes go.mod and go.sum with source code.
// go mod vendor	   ->    Creates a vendor directory with dependencies.
// go mod verify	   ->    Checks integrity of cached modules.
// go mod why	       ->    Explains why a package/module is required.
// go version -m	   ->    Reports Go version and embedded module information.
// go clean -modcache  ->    Removes the entire module cache.
