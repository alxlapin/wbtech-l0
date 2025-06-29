package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/order", getOrderHandler)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Request was handled")
}
