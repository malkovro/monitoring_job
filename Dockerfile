FROM golang:buster

RUN apt-get update && apt-get install -y gcc-arm-linux-gnueabi && rm -rf /var/lib/apt/lists/*
RUN go get -u github.com/d2r2/go-dht

WORKDIR /app

CMD CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build .