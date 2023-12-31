package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='color: #007BFF;'>Hello from Go Service! V5</h1>")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Go Service is running on :6011")
	http.ListenAndServe(":6011", nil)
}
