# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Create the final image
FROM alpine:latest


# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port that the application will run on
EXPOSE 8080

# Run the Go application
CMD ["./myapp"]