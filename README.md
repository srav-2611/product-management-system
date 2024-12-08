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
- Git installed to clone the repository.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/srav-2611/product-management-system.git
   cd product-management-system

