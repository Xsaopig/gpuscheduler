# Dockerfile
FROM golang:1.23 as builder
WORKDIR /app
COPY ./src/ .
RUN go mod tidy && go build -o main

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
