# Dockerfile
# Stage 1: Build
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o main ./cmd/main.go

# Stage 2: Production
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache tini ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder --chown=1000:1000 /app/main /app/main

# Create non-root user
RUN addgroup -g 1001 -S appuser && \
    adduser -S appuser -u 1001 && \
    chown -R appuser:appuser /app

USER appuser

# Expose port (fasthttp default)
EXPOSE 8080

# Health check
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#  CMD wget -qO- http://localhost:8080/health || exit 1

# Use tini for better signal handling
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/main"]