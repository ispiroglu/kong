package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Cihan!")
}
