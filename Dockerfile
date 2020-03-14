FROM golang:1.13.1
WORKDIR /go/src/github.com/kris-nova/tz0rk/

RUN go get github.com/kris-nova/logger
RUN go get github.com/spf13/cobra

COPY ./ ./

RUN go build -o /tz0rk .

ENTRYPOINT ["/tz0rk"]