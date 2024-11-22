FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main .

EXPOSE 3000

CMD [ "./main" ]

