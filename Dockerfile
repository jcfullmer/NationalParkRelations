# --- Stage 1: Build ---
FROM golang:1.25.6-alpine AS builder

# Install build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Run ---
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/main .

# IMPORTANT: Copy your static files into the final image
COPY --from=builder /app/static ./static

# Expose the port your server runs on
EXPOSE 8080

# Run the app
CMD ["./main"]
