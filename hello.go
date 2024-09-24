package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":8888", router)
}
