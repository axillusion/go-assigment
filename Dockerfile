# Use the official Go image.
FROM golang:1.23

WORKDIR /app

# Copy the Go module files first and download dependencies for caching
COPY go.mod go.sum ./
RUN go mod download

# Set the Go proxy environment variable
ENV GOPROXY=direct

# Copy the rest of the application code
COPY . .

# Build the Go app (we use 'main' for the output name)
RUN go build -o /main .

# The command that runs when the container starts
CMD ["/main"]