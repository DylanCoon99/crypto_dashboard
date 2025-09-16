# Stage 1: Builder
FROM golang:1.24.5

WORKDIR /app
COPY . .

# Build the Go binary
RUN go mod download
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the server
CMD ["./main"]