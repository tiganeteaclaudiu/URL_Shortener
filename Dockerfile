
# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="test"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./cmd/main ./cmd

# Expose port 8080 to the outside world
EXPOSE ${SERVICE_PORT}

# Command to run the executable
CMD ["./cmd/main", "redis"]
