# Use Go base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum from the root directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project code
COPY . .

# Set the microservice directory as the working directory
WORKDIR /app/microservices

# Build the image processor microservice
RUN go build -o image_processor .

# Start the image processor
CMD ["./image_processor"]
