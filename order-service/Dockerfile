# Stage 1: Build stage
FROM golang:1.20 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Final stage
FROM alpine:latest

# Install sqlite3 command-line tool
RUN apk add --no-cache sqlite

WORKDIR /root/

# Copy the built executable from builder stage
COPY --from=builder /app/main .

# Copy the database file from the current directory to the container
COPY test.db .

# Expose port 8082
EXPOSE 8082

# Command to run the executable
CMD ["./main"]
