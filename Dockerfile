# --- Build Stage ---
FROM golang:latest AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set workdir inside the container
WORKDIR /app

# Copy go mod and sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o bin/app cmd/*.go


# --- Final Stage ---
FROM alpine:latest

# Create a directory for the binary
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/bin/app .

# Set entrypoint
ENTRYPOINT ["./app"]
