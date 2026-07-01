# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
# Copy the binary from the builder
COPY --from=builder /app/main .
# Copy the jar file
COPY tnoodle.jar .
# Expose the port
EXPOSE 8080
# Run the application
CMD ["./main"]
