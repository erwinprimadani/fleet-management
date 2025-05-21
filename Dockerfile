# Build stage
FROM golang:1.24.2-alpine AS build

WORKDIR /

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /fleet-management ./cmd/main.go

# Final stage
FROM alpine:latest

# Add ca certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /

# Copy the binaries from the build stage
COPY --from=build /fleet-management /

# Create a non-root user to run the app
RUN adduser -D -g '' appuser
USER appuser

# Command to run
CMD ["/fleet-management"]