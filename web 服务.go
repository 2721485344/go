package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// each request calls handler
	log.Fatal(http.ListenAndServe("localhost:9512", nil)) }
// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你是我的眼")
}
