FROM golang:1.15-alpine

ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY golang/. .

RUN go build main.go

CMD ["./main"]