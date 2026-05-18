package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func logRequest(r *http.Request) {
	client := r.RemoteAddr
	log.Printf("%s, %s, %s", r.Method, r.RequestURI, client)
}

func home(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	secretsFile := os.Getenv("SECRET_FILE")
	var secret []byte
	var err error
	if secretsFile != "" {
		secret, err = os.ReadFile(secretsFile)
		if err != nil {
			log.Printf("error reading SECRET_FILE: %v", err)
		}
	}
	responses := []string{
		fmt.Sprintf("<p>Hello '%s' on '%s'!</p>", r.RemoteAddr, html.EscapeString(r.URL.Path)),
		fmt.Sprintf("<p>SECRET=%s</p>", html.EscapeString(os.Getenv("SECRET"))),
		fmt.Sprintf("<p>Secret from SECRET_FILE=%s</p>", html.EscapeString(string(secret))),
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<html><body>"+strings.Join(responses, "\n")+"</body></html>")
}

func main() {
	// Port to listen on
	port := 8888
	portRaw := os.Getenv("PORT")
	if portRaw != "" {
		var err error
		port, err = strconv.Atoi(portRaw)
		if err != nil {
			log.Fatalf("Invalid port: %v", err)
		}
	}

	http.HandleFunc("/", home)

	// Start the server
	serveStr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s\n", serveStr)
	log.Fatal(http.ListenAndServe(serveStr, nil))
}
