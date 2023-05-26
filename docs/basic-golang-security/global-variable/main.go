package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var insecureVariable string

func main() {
	http.HandleFunc("/api/change_variable/", handleVariableChange)
	http.HandleFunc("/", handleDefault)

	fmt.Println("Starting insecure web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleVariableChange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the variable content from the URL path
	path := r.URL.Path[len("/api/change_variable/"):]
	insecureVariable = strings.TrimPrefix(path, "/")

	fmt.Fprintf(w, "Variable changed to: %s", insecureVariable)
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insecure Variable: %s", insecureVariable)
}
