# Use the official Golang image as the base image
FROM golang:1.21.0 as build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and Go sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies specified in the go.mod and go.sum files
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-calorie-tracker

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the final image
COPY --from=build /app/go-calorie-tracker .

# Expose the port on which your application will run
EXPOSE 8000

# Command to run your application
CMD ["./go-calorie-tracker"]
