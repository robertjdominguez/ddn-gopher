FROM golang:latest

RUN go install github.com/air-verse/air@latest

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY auth-server/go.mod auth-server/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY auth-server/ .

# Run the application with air
CMD ["air"]
