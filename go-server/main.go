package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles a POST request to the "/form" endpoint.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request.
	if err := r.ParseForm(); err != nil {
		// If there's an error parsing the form, display the error message in the response.
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Display a success message in the response.
	fmt.Fprintf(w, "POST request successful\n")

	// Retrieve the values of the "name" and "address" fields from the form.
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Display the name and address in the response.
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler handles a GET request to the "/hello" endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/hello".
	if r.URL.Path != "/hello" {
		// If it's not "/hello", return a 404 Not Found error.
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the request method is GET.
	if r.Method != "GET" {
		// If it's not a GET request, return a 405 Method Not Allowed error.
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// If all checks pass, display "hello!" in the response.
	fmt.Fprint(w, "hello!")
}

func main() {
	// Create a file server that serves static files from the "./static" directory.
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle the root URL path ("/") with the file server.
	http.Handle("/", fileServer)

	// Register the formHandler function to handle requests to the "/form" endpoint.
	http.HandleFunc("/form", formHandler)

	// Register the helloHandler function to handle requests to the "/hello" endpoint.
	http.HandleFunc("/hello", helloHandler)

	// Print a message indicating that the server is starting.
	fmt.Printf("Starting server at port 8080\n")

	// Start the server and listen for incoming requests on port 8080.
	// If an error occurs, log the error and exit the program.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
