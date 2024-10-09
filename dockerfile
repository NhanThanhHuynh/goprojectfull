ARG GOARCH=amd64

# Stage 1: Build the Go application
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app
RUN go build -o main . && chmod +x ./main

# Check the built binary
RUN ls -l main

# Stage 2: Create a smaller image for production
FROM alpine:latest

# Install required packages
RUN apk add --no-cache ca-certificates

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Check the copied binary
RUN ls -l main

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]