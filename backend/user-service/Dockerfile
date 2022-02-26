FROM golang:1.16.6-alpine

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

CMD ["sh", "/app/main"]