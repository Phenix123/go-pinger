# Use a lightweight Go base image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy Go modules first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the application from the correct entry point in cmd/
RUN go build -o main ./cmd/main.go

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file (optional)
COPY .env .env

# Run the application
CMD ["./main"]
