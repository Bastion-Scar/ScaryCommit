# Используем официальный Go образ для сборки
FROM golang:1.24.4-alpine AS builder

# Устанавливаем зависимости для сборки (если нужны)
RUN apk add --no-cache git

# Копируем исходники
WORKDIR /app
COPY . .

# Собираем статический бинарник
RUN go build -o sco .

# Используем минимальный образ для финального бинарника
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Копируем бинарник из стадии сборки
COPY --from=builder /app/sco /usr/local/bin/sco

# Указываем точку входа (необязательно)
ENTRYPOINT ["sco"]
