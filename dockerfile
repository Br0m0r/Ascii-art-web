# First stage: builder
FROM golang:1.23.1-alpine AS builder

# Add metadata labels
LABEL version="1.0"
LABEL description="A web-based application to generate ASCII art in three banner styles: Standard, Shadow, and Thinkertoy."
LABEL maintainer="mfoteino, gpapadopoulos, vtsoucha"
LABEL project="ascii-art-web"

# Install security updates and ca-certificates
RUN apk update && apk add --no-cache ca-certificates

# Set the working directory for the build
WORKDIR /build

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY *.go ./
COPY templates/ templates/
COPY banners/ banners/

# Build the application with security flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o ascii-art-web .

# Second stage: final image
FROM alpine:latest

# Add metadata to final image
LABEL version="1.0"
LABEL description="ASCII Art Web Application - Production Image"
LABEL maintainer="mfoteino, gpapadopoulos, vtsoucha"

# Install security updates and create non-root user
RUN apk update && apk add --no-cache ca-certificates wget && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Create app directory and set ownership
WORKDIR /app
RUN chown appuser:appgroup /app

# Copy the binary from builder stage
COPY --from=builder --chown=appuser:appgroup /build/ascii-art-web .

# Copy templates and banners directories
COPY --from=builder --chown=appuser:appgroup /build/templates/ templates/
COPY --from=builder --chown=appuser:appgroup /build/banners/ banners/

# Switch to non-root user
USER appuser

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./ascii-art-web"]

