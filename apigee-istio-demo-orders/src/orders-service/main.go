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
	server.HandleFunc("/orders", orders)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}


func orders(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"Status\":\"Processing...\",")
    fmt.Fprintf(w, "\"host\":\""+host+"\",")
    fmt.Fprintf(w, "\"Expected Delivery\":\"" +strconv.Itoa(rand.Intn(25))+ " Minutes\"}")
	
}
// [END all]