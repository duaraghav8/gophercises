package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/duaraghav8/gophercises/ex2_urlshort/urlshort"
	"github.com/gorilla/mux"
)

func main() {
	const jsonPathsFile = "./url-mappings.json"
	router := mux.NewRouter()
	paths := map[string]string{
		"/me_linkedin": "https://www.linkedin.com/in/raghavdua/",
		"/solium_npm":  "https://www.npmjs.com/package/solium",
	}
	jsonPaths, err := ioutil.ReadFile(jsonPathsFile)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error occured while opening %s: %s", jsonPathsFile, err))
		panic(err)
	}

	mapHandler := urlshort.MapHandler(paths, router)
	jsonHandler, err := urlshort.JSONHandler(jsonPaths, mapHandler)

	if err != nil {
		log.Fatal("There was a problem configuring mux using json handler: ", err)
		panic(err)
	}

	log.Println("Started listening on http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", jsonHandler)
}
