# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run the binary.
CMD ["/app/main"]