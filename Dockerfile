FROM golang:1.20

WORKDIR /app

ADD . /app

RUN go mod tidy

CMD go test -v
