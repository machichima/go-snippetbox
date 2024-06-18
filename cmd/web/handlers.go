package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// def home handler
// Exec logic
func home(w http.ResponseWriter, r *http.Request) {
	// check whether the url is exactly "/", as "/" is a subtree path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello"))
}

// view all snippet
func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// create snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// Need  to define all the header content before WriteHeader function!
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed!"))
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create new snippet..."))
}
