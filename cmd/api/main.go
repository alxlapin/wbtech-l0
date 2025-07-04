package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"wbtech-l0/internal/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var orderService *service.OrderService

func main() {
	dsn, _ := os.LookupEnv("GORM_DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	orderService = &service.OrderService{GormDB: db}

	http.HandleFunc("GET /order/{id}", getOrderHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.ListenAndServe(":8080", nil)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.PathValue("id")

	order, err := orderService.GetOrderByID(orderID)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(order)
	}
}
