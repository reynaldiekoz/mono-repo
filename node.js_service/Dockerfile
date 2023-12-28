# Use an official Node.js runtime as a parent image
FROM node:14-alpine

# Set the working directory to /nodejs_service
WORKDIR /nodejs_service

# Copy the current directory contents into the container at /nodejs_service
COPY . .

# Install dependencies
RUN npm install

# Expose port 6012 to the outside world
EXPOSE 6012

# Command to run the Node.js service
CMD ["node", "server.js"]
