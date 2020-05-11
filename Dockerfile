FROM golang:latest

WORKDIR /go/src/shop

COPY ./supermarket/market.go /supermarket/market.go

CMD ["go","run","main.go"]
