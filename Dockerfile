# Step 1: Build Go Application
FROM golang:1.21-alpine AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gocas ./cmd

FROM alpine:latest

WORKDIR /app

# Copy Go application binary
COPY --from=go-builder /app/gocas .


# Expose port
EXPOSE 8080

# Set the entry point
CMD ["./gocas","-app","rest"]
