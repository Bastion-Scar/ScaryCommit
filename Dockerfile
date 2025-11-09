FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o sco-linux .
RUN GOOS=windows GOARCH=amd64 go build -o sco-windows.exe .
RUN GOOS=darwin GOARCH=amd64 go build -o sco-macos .

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/sco-linux /usr/local/bin/sco-linux
COPY --from=builder /app/sco-windows.exe /usr/local/bin/sco-windows.exe
COPY --from=builder /app/sco-macos /usr/local/bin/sco-macos

RUN ln -s /usr/local/bin/sco-linux /usr/local/bin/sco