package main

import (
	"fmt"
	"log"
	"net/http"
)

// handle the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check if URL path is correct
	if r.URL.Path != "/hello" {
		// send 404 error response
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// only allow GET method calls
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// handle /form route
func formHandler(w http.ResponseWriter, r *http.Request) {
	// check if error in form parsing
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	// fetch form data values
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
}

func main() {
	// initialize file server
	fileServer := http.FileServer(http.Dir("./static"))

	// HTTP handlers
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	// start the web server and check for errors, if any
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
