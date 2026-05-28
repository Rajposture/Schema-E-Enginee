# Use official Go image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the app
RUN go build -o main ./cmd/server

# Run the app
CMD ["./main"]