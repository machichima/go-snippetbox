package main

import (
	"log"
	"net/http"
)

func main() {
	// init servemux (router)
	// register home function as the handler for root "/"
	mux := http.NewServeMux()
	mux.HandleFunc("/", home) // route, function
	mux.HandleFunc("/snippets", viewSnippet)
	mux.HandleFunc("/snippets/post", createSnippet)

	// start a new web server
	var port string = "4000"
	log.Printf("Starting server on :%s", port)

	// convert port to string
	err_http := http.ListenAndServe(":"+port, mux)
	log.Fatal(err_http)
}
