##############
## Build image
##############
FROM golang:1.12-alpine as build
LABEL maintainer="cx@initd.cz"

ENV GOPATH=/srv/go
RUN apk add --no-cache make git gcc libc-dev

WORKDIR /srv
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /srv/go/src/github/by-cx/concourse-build-number-resource
RUN export PATH=$PATH:/srv/go/bin && cd /srv/go/src/github/by-cx/concourse-build-number-resource && make

############
## App image
############

FROM alpine:3.9
LABEL maintainer="cx@initd.cz"

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build /srv/go/src/github/by-cx/concourse-build-number-resource/bin/in /opt/resource/
COPY --from=build /srv/go/src/github/by-cx/concourse-build-number-resource/bin/check /opt/resource/
COPY --from=build /srv/go/src/github/by-cx/concourse-build-number-resource/bin/out /opt/resource/