package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /order/{id}", getOrderHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.PathValue("id")

	fmt.Fprintf(w, "Request was handled ", orderID)
}
