FROM golang:1.21.0

WORKDIR /app

# Copy files and download dependencies

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Define build arguments for sensitive data
# These values will be provided during container runtime
ARG DATABASE_URL

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-calorie-tracker

# Expose the port on which your application will run
EXPOSE 8000

# Define environment variables with sensitive data
# These values will be provided during container runtime
ENV DATABASE_URL=${DATABASE_URL}

# Command to run your application
CMD ["./go-calorie-tracker"]