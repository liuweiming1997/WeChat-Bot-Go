FROM golang:1.13.1-alpine3.10

COPY . $GOPATH/src/github.com/WeChat-Bot-Go

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR $GOPATH/src/github.com/WeChat-Bot-Go

RUN go build -o main main.go

ENTRYPOINT ["./main"]
