# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /build

# CHANGE: Remove 'app/' because go.mod is in the current directory
COPY go.mod ./

RUN go mod download

# CHANGE: Remove 'app/' because main.go is in the current directory
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/main .
EXPOSE 8080
CMD ["./main"]
