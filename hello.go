package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World1")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hell", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":8888", router)
}
