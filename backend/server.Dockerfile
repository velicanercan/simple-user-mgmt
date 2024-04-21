# Build Stage
FROM golang:latest AS builder

WORKDIR /go/src/app

COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o simple-user-mgmt main.go

# Final Stage
FROM gcr.io/distroless/base-debian10

COPY --from=builder /go/src/app/simple-user-mgmt /

EXPOSE ${SERVER_PORT}

ENTRYPOINT ["/simple-user-mgmt"]
