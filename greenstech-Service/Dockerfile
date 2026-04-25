# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o app main.go


# ---------- Stage 2: Run ----------
FROM alpine:latest

# Install CA certificates (for HTTPS calls)
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/app .

# Copy .env if needed (optional)


# Expose port (change if your app uses different port)
EXPOSE 8080

# Run the app
CMD ["./app"]
