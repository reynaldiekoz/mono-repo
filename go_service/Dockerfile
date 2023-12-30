# === Stage 1: Build Stage ===
FROM golang:1.19-alpine as builder

# Set the working directory to /go_service
WORKDIR /go_service

# Copy the application code
COPY . .

# Build the Go service
RUN CGO_ENABLED=0 GOOS=linux go build -o go-service

# === Stage 2: Final Stage ===
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /go_service/go-service .

# Expose port 6011 to the outside world
EXPOSE 6011

# Command to run the executable
CMD ["./go-service"]