# Use an official Go runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory to /go_service
WORKDIR /go_service

# Copy the go.mod file into the container
COPY go.mod .

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go service
RUN go build -o go-service

# Expose port 6012 to the outside world
EXPOSE 6011

# Command to run the executable
CMD ["./go-service"]
