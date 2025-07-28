package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleFunc)

	fmt.Println("Running Servre at 8080")
	http.ListenAndServe(":8080", nil)
}
