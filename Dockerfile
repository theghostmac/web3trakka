# Use a Go base image.
FROM golang:1.21.4-bookworm as builder

# Set the workig directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the application
RUN go build -o web3trakka ./cmd/main.go

# Use a smaller base image for the final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/web3trakka .

# Expose the port the app runs on
EXPOSE 7080

# Run the binary
CMD ["./web3trakka"]
