# Use Go 1.24 (or higher)
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
# Copy nested module's go.mod/go.sum so replace directive resolves
COPY getOutletFood/go.mod getOutletFood/go.sum ./getOutletFood/

# Copy the rest of the source code
COPY . .

# Download dependencies (now that all go.mod files exist)
RUN go mod download

# Build the Go binary
RUN go build -o myservice .

# Final image (small, just the binary)
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/myservice .

CMD ["./myservice"]
