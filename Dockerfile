# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11

# Copy the binary from the builder
COPY --from=builder /main /main

# Expose the application port
EXPOSE 8080

# Set the entrypoint to the binary
ENTRYPOINT ["/main"]