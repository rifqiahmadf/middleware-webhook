# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o middleware-webhook main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary file
COPY --from=builder /app/middleware-webhook .

# Expose port
EXPOSE 80

# Command to run the executable
CMD ["./middleware-webhook"]