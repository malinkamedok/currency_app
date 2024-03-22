# Шаг 1: Установка зависимостей
FROM golang:1.22-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod tidy

# Step 2: Сборка проекта
FROM golang:1.22-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/main

# Step 3: Запуск
FROM scratch
EXPOSE 8080
COPY --from=builder /bin/app /usr/bin/app
CMD [ "/usr/bin/app" ]