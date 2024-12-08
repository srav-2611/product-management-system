# Product Management System

A modular and scalable product management system built with Go, PostgreSQL, Redis, RabbitMQ, and Docker. This system demonstrates modern backend development practices, including microservices architecture, caching strategies, and message queue-based task processing.

---

## Features

- **CRUD Operations**: Manage products with endpoints for creating, reading, updating, and deleting.
- **Redis Caching**: Optimized read performance using Redis for frequently accessed data.
- **Image Processing**: Asynchronous image compression using RabbitMQ and a microservice-based architecture.
- **Database Integration**: Persistent storage with PostgreSQL.
- **Containerization**: Fully containerized using Docker and Docker Compose for seamless deployment.

---

## Architecture Overview

1. **API Service**:
   - Handles product management endpoints.
   - Uses Redis for caching product data.
   - Communicates with PostgreSQL for persistent storage.

2. **Image Processing Microservice**:
   - Listens to RabbitMQ for image processing tasks.
   - Updates the database with compressed image URLs.

3. **PostgreSQL**:
   - Stores product information in a structured schema.

4. **Redis**:
   - Speeds up read operations by caching product details.

5. **RabbitMQ**:
   - Enables asynchronous task processing for image compression.

---

## Setup Instructions

### Prerequisites
- Docker and Docker Compose installed.
- Git was installed to clone the repository.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/srav-2611/product-management-system.git
   cd product-management-system
2. Build and start the application:

bash
Copy code
docker-compose up --build

3. Access the API:
Base URL: http://localhost:8080

API Endpoints
Products
GET /products/{id}: Fetch product details by ID (with Redis caching).
POST /products: Create a new product.
PUT /products/{id}: Update an existing product.
DELETE /products/{id}: Delete a product and invalidate its cache.
Image Processing
Automatically processes product images and compresses them.

Configuration
Database
-POSTGRES_USER: Postgres
-POSTGRES_PASSWORD: Comrade#11
-POSTGRES_DB: product_management

Redis
- Configured to run on default port 6379.
RabbitMQ
- Accessible via localhost:5672.

Assumptions
1. Product IDs are unique and auto-incremented.
2 . All endpoints are tested locally on http://localhost:8080.
3. Redis cache expiration is set to 10 minutes.
   
Testing
- Unit and integration tests cover major functionalities.
- Use Postman or Curl to test API endpoints.

Example Usage
Add a New Product
- Endpoint: POST /products
- Request Body:
  {
  "product_name": "Sample Product",
  "product_description": "A sample product description",
  "product_price": 100.0,
  "product_images": ["http://example.com/image1.jpg"]
}
- Expected Response:
  {
  "message": "Product created successfully",
  "product_id": 1
}

Fetch a Product
- Endpoint: GET /products/1
- Expected Response:
  {
  "id": 1,
  "user_id": 1,
  "product_name": "Sample Product",
  "product_description": "A sample product description",
  "product_images": ["http://example.com/image1.jpg"],
  "compressed_product_images": ["http://example.com/image1-compressed.jpg"],
  "product_price": 100.0
}

Deployment Notes
1. Ensure all environment variables (e.g., database credentials, Redis URL, RabbitMQ URL) are correctly set in the docker-compose.yml file.
2. Use a production-ready database and Redis setup for live environments.
3. Set up monitoring for RabbitMQ queues and database performance for optimized scaling.
