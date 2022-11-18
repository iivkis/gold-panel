FROM golang:alpine

WORKDIR /app

COPY . .

CMD ["go", "run", "./cmd/panel/main.go"]