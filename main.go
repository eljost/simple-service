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
	log.Printf("%s, %s, %s", r.Method, r.RequestURI, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
