
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod init password-cracker && \
    go mod tidy && \
    go build -o password-cracker .


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/password-cracker .
COPY database.json .
ENTRYPOINT ["./password-cracker"]