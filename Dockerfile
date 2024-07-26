FROM golang:1.22.5-alpine

# Installing dependencies
RUN apk update && apk add --no-cache make postgresql-client
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .
# Copy the start script
COPY start.sh /start.sh
RUN chmod +x /start.sh

# Expose the port that the application will run on
EXPOSE ${PORT}

# Set the entrypoint command to run the application
CMD ["/start.sh"]
