FROM golang:1.24.4-alpine AS builder

# Install depends
RUN apk add --no-cache git

# Copy app
WORKDIR /app
COPY . .

# Building binary
RUN go build -o sco .

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/sco /usr/local/bin/sco

ENTRYPOINT ["sco"]
