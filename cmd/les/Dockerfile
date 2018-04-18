FROM golang:latest AS build

# without this 'go get' fails with "go get: no install location for directory /go/src outside GOPATH"
ENV GOBIN /go/bin

WORKDIR /go/src

COPY *.go ./

RUN go get

# need to statically link, as `alpine` (and `scratch`) don't have libc
RUN CGO_ENABLED=0 go build -o ../bin/les


# ---------------------------

# could use `scratch` but this lets you `sh` to help diagnose image issues
FROM alpine:latest

RUN mkdir -p /go/bin
COPY --from=build /go/bin/les /go/bin/

WORKDIR /les

RUN addgroup -g 1000 les \
 && adduser -u 1000 -G les -s /bin/sh -D les

USER les

ENTRYPOINT ["/go/bin/les"]

