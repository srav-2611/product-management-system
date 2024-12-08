package main

import (
	"fmt"
	"log"
	"net/http"
	"product-management-system/config"
	"product-management-system/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database and Redis
	config.ConnectDatabase()
	defer config.CloseDatabase()
	config.ConnectRedis()
	defer config.CloseRedis()

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/products/{id}", handlers.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.DeleteProductHandler).Methods("DELETE")
	router.HandleFunc("/products/{id}", handlers.UpdateProductHandler).Methods("PUT")

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
