package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World1")
}

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(helloHandler))
}
