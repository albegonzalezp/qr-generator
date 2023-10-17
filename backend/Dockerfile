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

# Set environment variables
ENV HOST=smtp.gmail.com
ENV PORT=587
ENV FROM=qrgenerator0@gmail.com
ENV PASSWORD=sbknrzlqqsetztkq
ENV SUBJECT=": by zetacoder"
ENV BODY="Hi! Here is your <b>QR code</b> attached for the link you provided.<br><br>Regards,<br>ZetaCoder"
ENV TEMPLATES_DIR=C:\\Users\\zalbe\\Desktop\\go-workspace\\src\\github.com\\zetacoder\\personal\\qr-generator\\frontend\\templates
ENV STATIC_DIR=C:\\Users\\zalbe\\Desktop\\go-workspace\\src\\github.com\\zetacoder\\personal\\qr-generator\\frontend\\assets
ENV IMAGES_DIR=C:\\Users\\zalbe\\Desktop\\go-workspace\\src\\github.com\\zetacoder\\personal\\qr-generator\\backend\\qr_images


# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port that the application will run on
EXPOSE 8080

# Run the Go application
CMD ["./myapp"]