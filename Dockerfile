FROM golang

ADD . /go/src/github.com/deaswang/goproc
RUN go install github.com/deaswang/goproc

ENTRYPOINT /go/bin/goproc

EXPOSE 3000