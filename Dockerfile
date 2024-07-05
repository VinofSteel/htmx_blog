FROM golang:1.22.3-alpine

# Installing dependencies
RUN apk update && apk add --no-cache make
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

# Build the application
# RUN make build

# Expose the port that the application will run on
EXPOSE ${PORT}

# Set the entrypoint command to run the application
CMD ["sh", "-c", "if [ \"$ENV\" = \"production\" ]; then make build && ./htmx_blog; else make run; fi"]
