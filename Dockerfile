# Stage 1: Build the Go application
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files for dependencies
COPY go.mod go.sum ./

# Download all Go modules (dependencies)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o /app/main .

# Stage 2: Run the application with a minimal base image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage to the final container
COPY --from=builder /app/main .

# Copy the .env file if necessary (optional)
COPY --from=builder /app/.env .
COPY --from=builder /app/static ./static

# Expose the port that the app listens on
EXPOSE 3000

# Set the entry point to the executable
CMD ["./main"]
