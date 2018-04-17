FROM golang:latest AS build

# without this 'go get' fails with "go get: no install location for directory /go/src outside GOPATH"
ENV GOBIN /go/bin

WORKDIR /go/src

COPY *.go ./

RUN go get

# because of DNS lookups and `os/user` it's quite complicated to do a statically-linked binary :(
RUN go build -o ../bin/les-node


# ---------------------------

# if there's a way to support `libc` and `os/user` from `scratch` or `alpine` I haven't found it :(
FROM ubuntu:latest

# this avoids "x509: failed to load system roots and no roots provided"
RUN apt-get update
RUN apt-get install -y ca-certificates

RUN mkdir -p /go/bin
COPY --from=build /go/bin/les-node /go/bin/

WORKDIR /les

RUN groupadd --gid 1000 les \
  && useradd --uid 1000 --gid les --shell /bin/bash --create-home les

USER les

ENTRYPOINT ["/go/bin/les-node"]

