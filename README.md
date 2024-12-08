#Product Management System
A modular and scalable product management system built with Go, PostgreSQL, Redis, RabbitMQ, and Docker. This system demonstrates modern backend development practices, including microservices architecture, caching strategies, and message queue-based task processing.

Features
CRUD Operations: Manage products with endpoints for creating, reading, updating, and deleting.
Redis Caching: Optimized read performance using Redis for frequently accessed data.
Image Processing: Asynchronous image compression using RabbitMQ and a microservice-based architecture.
Database Integration: Persistent storage with PostgreSQL.
Containerization: Fully containerized using Docker and Docker Compose for seamless deployment.
Architecture Overview
API Service:
Handles product management endpoints.
Uses Redis for caching product data.
Communicates with PostgreSQL for persistent storage.
Image Processing Microservice:
Listens to RabbitMQ for image processing tasks.
Updates the database with compressed image URLs.
PostgreSQL:
Stores product information in a structured schema.
Redis:
Speeds up read operations by caching product details.
RabbitMQ:
Enables asynchronous task processing for image compression.
Setup Instructions
Prerequisites
Docker and Docker Compose installed.
Git installed to clone the repository.
Steps
Clone the repository:
bash
Copy code
git clone https://github.com/srav-2611/product-management-system.git
cd product-management-system
Build and start the application:
bash
Copy code
docker-compose up --build
Access the API:
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
Database:
POSTGRES_USER: postgres
POSTGRES_PASSWORD: Comrade#11
POSTGRES_DB: product_management
Redis: Configured to run on default port 6379.
RabbitMQ: Accessible via localhost:5672.
Assumptions
Product IDs are unique and auto-incremented.
All endpoints are tested locally on http://localhost:8080.
Redis cache expiration is set to 10 minutes.
Testing
Unit and integration tests cover major functionalities.
Use Postman or curl to test API endpoints.
