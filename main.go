package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main() {
	fmt.Println("Initialising URL shortener")

	// Get JSON from file and turn it into a map
	var redirects map[string]string
	file, err := ioutil.ReadFile("redirects.json")
	if err != nil { log.Fatal(err) }
	err = json.Unmarshal(file, &redirects)
	if err != nil { log.Fatal(err) }
	fmt.Println(redirects)

	// Handle all requests
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// Check if redirect exists
		var path = r.URL.Path
		_, exists := redirects[path]

		// Redirect
		if exists == true {
			http.Redirect(w, r, redirects[path], 303)
		} else {
			http.Redirect(w, r, redirects["/"], 404)
		}
	})
	
	// Listen
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	fmt.Println("Listening on " + port)
	http.ListenAndServe(":" + port, nil)
}