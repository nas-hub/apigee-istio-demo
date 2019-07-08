// [START all]
package main

import (
	"fmt"
	"log"
    "strconv"
	"net/http"
	"os"
    "math/rand"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "5000"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/inventory", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"list\":[{\"Status\": \"OK\"},")
    fmt.Fprintf(w, "{\"Product One\":\"" +strconv.Itoa(rand.Intn(100))+ "\"},")
    fmt.Fprintf(w, "{\"Product Two\":\"" +strconv.Itoa(rand.Intn(100))+ "\"},")
    fmt.Fprintf(w, "{\"Product Three\":\"" +strconv.Itoa(rand.Intn(100))+ "\"},")
    fmt.Fprintf(w, "{\"Product Four\":\"" +strconv.Itoa(rand.Intn(100))+ "\"}]}")
}
// [END all]