package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product-management-system/config"
	"product-management-system/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	// Define cache key
	cacheKey := fmt.Sprintf("product:%d", id)

	// Check Redis cache
	cachedProduct, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		// Cache hit
		var product models.Product
		if err := json.Unmarshal([]byte(cachedProduct), &product); err == nil {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// Cache miss: fetch from database
	product, err := models.GetProductByID(id)
	if err != nil {
		http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
		return
	}

	// Save to Redis
	productJSON, _ := json.Marshal(product)
	config.RedisClient.Set(config.Ctx, cacheKey, productJSON, 10*time.Minute)

	json.NewEncoder(w).Encode(product)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Insert the product into the database
	productID, err := models.CreateProduct(product)
	if err != nil {
		http.Error(w, `{"error":"Failed to create product"}`, http.StatusInternalServerError)
		return
	}

	// Return success response with the created product ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Product created successfully",
		"product_id": productID,
	})
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	err = models.DeleteProduct(id)
	if err != nil {
		http.Error(w, `{"error":"Failed to delete product"}`, http.StatusInternalServerError)
		return
	}

	// Invalidate Redis cache
	cacheKey := fmt.Sprintf("product:%d", id)
	config.RedisClient.Del(config.Ctx, cacheKey)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	err = models.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, `{"error":"Failed to update product"}`, http.StatusInternalServerError)
		return
	}

	// Invalidate Redis cache
	cacheKey := fmt.Sprintf("product:%d", id)
	config.RedisClient.Del(config.Ctx, cacheKey)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}
