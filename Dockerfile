FROM golang:1.12-alpine
LABEL maintainer="cx@initd.cz"

ENV GOPATH=/srv/go
RUN apk add --no-cache make git

WORKDIR /srv
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /srv/go/src/github/by-cx/concourse-build-number-resource
RUN export PATH=$PATH:/srv/go/bin && cd /srv/go/src/github/by-cx/concourse-build-number-resource && make
