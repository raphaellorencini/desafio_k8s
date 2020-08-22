FROM golang:1.15rc1-alpine3.12

COPY . .

RUN go build main.go

ENTRYPOINT ["./main"]