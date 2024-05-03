# syntax=docker/dockerfile:1

FROM golang:1.22.1 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod ./

# Download all dependencies. 
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go app as a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-accelone ./cmd

# Start from a smaller image for a smaller final image.
FROM alpine:latest  

# Set the Current Working Directory inside the container.
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/go-accelone .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./go-accelone"]