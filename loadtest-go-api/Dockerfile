# Dockerfile for Go API

# Use the official Golang image as a build stage
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o loadtest-go-api main.go

# Use a smaller base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/loadtest-go-api .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./loadtest-go-api"]
