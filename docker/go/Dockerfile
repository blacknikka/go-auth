FROM golang:1.13.4-buster

WORKDIR /go/src/github.com/blacknikka/go-auth

RUN go get -u github.com/go-delve/delve/cmd/dlv \
    && go get -u github.com/oxequa/realize \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/go-sql-driver/mysql \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.7.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv ./migrate.linux-amd64 /usr/bin/

ENV PATH $PATH:/go/bin

CMD ["go", "run", "./main.go"]
