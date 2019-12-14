FROM golang:1.13.4-buster

WORKDIR /go

RUN go get -u github.com/go-delve/delve/cmd/dlv

ENV PATH $PATH:/go/bin

CMD ["go", "run", "./app/main.go"]
