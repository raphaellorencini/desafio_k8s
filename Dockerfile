FROM golang:1.15rc1-alpine3.12

WORKDIR /go/src/app
COPY golang/. .

RUN go build main.go

ENTRYPOINT ["./main"]