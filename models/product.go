package models

import (
	"context"
	"product-management-system/config"
)

type Product struct {
	ID                 int      `json:"id"`
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	CompressedImages   []string `json:"compressed_product_images"`
	ProductPrice       float64  `json:"product_price"`
}

// GetProducts fetches products with optional filtering, sorting, pagination, and search
func GetProducts(userID, minPrice, maxPrice, limit, offset, sort, order, search string) ([]Product, error) {
	query := `
        SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price
        FROM products
        WHERE ($1::int IS NULL OR user_id = $1::int)
          AND ($2::numeric IS NULL OR product_price >= $2::numeric)
          AND ($3::numeric IS NULL OR product_price <= $3::numeric)
          AND ($4::text IS NULL OR product_name ILIKE '%' || $4 || '%')
        ORDER BY $5 $6
        LIMIT $7 OFFSET $8
    `

	rows, err := config.DB.Query(context.Background(), query,
		toNullableInt(userID),
		toNullableFloat(minPrice),
		toNullableFloat(maxPrice),
		toNullableString(search),
		sort,
		order,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.CompressedImages, &product.ProductPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductByID fetches a single product by its ID
func GetProductByID(id int) (Product, error) {
	query := `
        SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price
        FROM products
        WHERE id = $1
    `
	var product Product
	err := config.DB.QueryRow(context.Background(), query, id).Scan(
		&product.ID,
		&product.UserID,
		&product.ProductName,
		&product.ProductDescription,
		&product.ProductImages,
		&product.CompressedImages,
		&product.ProductPrice,
	)
	return product, err
}

// CreateProduct inserts a new product into the database
func CreateProduct(product Product) (int, error) {
	query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	var productID int
	err := config.DB.QueryRow(context.Background(), query,
		product.UserID,
		product.ProductName,
		product.ProductDescription,
		product.ProductImages,
		product.ProductPrice,
	).Scan(&productID)
	return productID, err
}

// DeleteProduct deletes a product by its ID
func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := config.DB.Exec(context.Background(), query, id)
	return err
}

// UpdateProduct updates a product in the database by its ID
func UpdateProduct(id int, product Product) error {
	query := `
        UPDATE products
        SET product_name = $1,
            product_description = $2,
            product_images = $3,
            product_price = $4
        WHERE id = $5
    `
	_, err := config.DB.Exec(context.Background(), query,
		product.ProductName,
		product.ProductDescription,
		product.ProductImages,
		product.ProductPrice,
		id,
	)
	return err
}

// Helper functions for nullable parameters
func toNullableInt(value string) interface{} {
	if value == "" {
		return nil
	}
	return value
}

func toNullableFloat(value string) interface{} {
	if value == "" {
		return nil
	}
	return value
}

func toNullableString(value string) interface{} {
	if value == "" {
		return nil
	}
	return value
}
