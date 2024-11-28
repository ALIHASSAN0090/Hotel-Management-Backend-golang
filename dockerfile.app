FROM golang:1.23-bullseye

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 3000

CMD ["go", "run", "main.go"]