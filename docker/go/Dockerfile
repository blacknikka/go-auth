FROM golang:1.13.4-buster

WORKDIR /go/src/github.com/blacknikka/go-auth

RUN go get -u github.com/go-delve/delve/cmd/dlv \
    && go get github.com/oxequa/realize

ENV PATH $PATH:/go/bin

CMD ["go", "run", "./main.go"]
