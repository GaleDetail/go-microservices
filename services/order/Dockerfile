FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o order-service .

EXPOSE 50052

CMD ["./order-service"]
