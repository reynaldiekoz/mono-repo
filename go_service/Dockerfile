# Use an official Go runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory to /go_service
WORKDIR /go_service

# Copy the current directory contents into the container at /go_service
COPY . .

# Build the Go service
RUN go build -o go-service

# Expose port 6012 to the outside world
EXPOSE 6011

# Command to run the executable
CMD ["./go-service"]
