FROM golang:1.7.3

## RUN go get github.com/axw/gocov/gocov && go get -u gopkg.in/matm/v1/gocov-html

COPY . /go/src/file_share
COPY ./config_docker.toml /go/src/file_share/config.toml

WORKDIR /go/src/file_share

RUN go get && go build

CMD ./file_share