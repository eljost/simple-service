package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	client := r.RemoteAddr
	log.Printf("%s, %s, %s", r.Method, r.RequestURI, client)
	fmt.Fprintf(w, "Hello '%s' on '%q'!", client, html.EscapeString(r.URL.Path))
}

func main() {
	port_raw := os.Getenv("PORT")
	port, err := strconv.Atoi(port_raw)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	http.HandleFunc("/", home)

	serveStr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s\n", serveStr)
	log.Fatal(http.ListenAndServe(serveStr, nil))
}
