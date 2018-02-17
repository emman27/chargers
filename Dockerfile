FROM golang

WORKDIR /go/src/github.com/emman27/chargers

RUN go get -u github.com/pilu/fresh
RUN go get -u github.com/golang/dep/cmd/dep
RUN go install github.com/pilu/fresh

COPY api api
COPY constants constants
COPY controllers controllers
COPY db db
COPY main.go main.go
COPY Gopkg.lock Gopkg.lock
COPY Gopkg.toml Gopkg.toml

RUN dep ensure

EXPOSE 8080

CMD ["/go/bin/fresh"]
