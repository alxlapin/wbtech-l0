# Build stage
FROM golang:1.24.4-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o wbtechl0 .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/wbtechl0 /app/index.html ./

# Command to run the application
CMD ["./wbtechl0"]