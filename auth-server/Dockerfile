# syntax=docker/dockerfile:1

# Use the official Golang image to create a build artifact.
# This image includes the necessary Go tools.
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./auth-server/go.mod ./auth-server/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./auth-server .

# Build the Go app with build arguments
ARG JWT_SECRET
ENV JWT_SECRET=${JWT_SECRET}
RUN go build -o auth-server ./main.go

# Start a new stage from scratch
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/auth-server /auth-server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/auth-server"]
