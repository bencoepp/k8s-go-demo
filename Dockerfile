FROM golang:1.21-alpine

WORKDIR /app

# Copy app source
COPY go.mod ./
COPY main.go ./

# Build binary
RUN go mod tidy && go build -o demo-app

# Add config
RUN mkdir -p /etc/demo
COPY config/config.txt /etc/demo/config

# Set environment variable (can be overridden with `-e`)
ENV MESSAGE="Hello from Docker!"

# Expose port
EXPOSE 8080

CMD ["./demo-app"]
