# Use the official Golang image as a build stage
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Start an interactive shell instead of building the app
CMD ["/bin/sh"]
# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Install sqlite3 command-line tool
RUN apk add --no-cache sqlite

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the database file from the current directory to the container
COPY test.db .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
