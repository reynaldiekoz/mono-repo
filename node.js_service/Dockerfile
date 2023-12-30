# === Stage 1: Build Stage ===
FROM node:18-alpine as builder

# Set the working directory to /nodejs_service
WORKDIR /nodejs_service

# Copy the current directory contents into the container at /nodejs_service
COPY . .

# Install dependencies
RUN npm install  --only=production && npm prune

# === Stage 2: Final Stage ===
FROM node:18-alpine 

# Set the working directory to /app
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /nodejs_service .

# Expose port 6012 to the outside world
EXPOSE 6012

# Command to run the Node.js service
CMD ["node", "server.js"]

