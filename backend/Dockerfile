# Dockerfile
FROM golang:1.23 as builder
WORKDIR /app
COPY ./src/ .
RUN go mod tidy && go build -o main

# Final stage
FROM alpine:latest
WORKDIR /root/

# Install glibc
RUN apk --no-cache add \
    libc6-compat

COPY --from=builder /app/main .
RUN chmod +x ./main
EXPOSE 8080
CMD ["./main"]
