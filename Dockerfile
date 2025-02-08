FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Сборка приложения
RUN go build -v -o /usr/src/app/cmd/app/app ./cmd/app/main.go  # Создаем исполняемый файл с именем app в директории cmd/app

# Указываем, что делать при запуске контейнера
CMD ["./cmd/app/app"]