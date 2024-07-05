# Use the official Golang image as the base image
FROM golang:1.22.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -mod=vendor -o htmx_blog

# Expose the port that the application will run on
EXPOSE ${PORT}

# Set the entrypoint command to run the application
CMD ["./htmx_blog"]
