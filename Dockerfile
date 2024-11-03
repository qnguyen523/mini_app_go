# Stage 1: Build the Go binary
FROM golang:1.22.5-alpine AS builder
# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod ./
RUN go mod download
# Copy the rest of the application code
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
# Stage 2: Create a lightweight image for running the Go application
FROM alpine:latest
# Set the working directory
WORKDIR /root/
# Copy the built binary from the builder stage
COPY --from=builder /app/app .
# Expose the port the app runs on
EXPOSE 8080
# Run the binary
CMD ["./app"]
