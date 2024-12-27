# Dockerfile
FROM m.daocloud.io/docker.io/library/golang:1.23 as builder
WORKDIR /app
COPY ./src/ .
RUN go mod tidy && go build -o main

# Final stage
FROM m.daocloud.io/docker.io/library/alpine:latest
WORKDIR /root/

# Install glibc
RUN apk --no-cache add ca-certificates && \
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.35-r0/glibc-2.35-r0.apk && \
    apk add glibc-2.35-r0.apk && \
    rm glibc-2.35-r0.apk

COPY --from=builder /app/main .
RUN chmod +x ./main
EXPOSE 8080
CMD ["./main"]
