FROM golang:1.12-alpine

ADD . /go/src/github.com/oleewere/cldnt
WORKDIR /go/src/github.com/oleewere/cldnt
ENV GO111MODULE=on
RUN apk add --no-cache git
RUN go get -u github.com/gobuffalo/packr/packr
RUN packr -v
RUN go build -o /cldnt .

FROM alpine:3.7
RUN apk add --no-cache ca-certificates

COPY --from=0 /cldnt /

ENTRYPOINT ["/cldnt"]