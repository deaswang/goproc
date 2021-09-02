FROM golang
WORKDIR /go/src/

ADD . /go/src/github.com/deaswang/goproc
RUN cd /go/src/github.com/deaswang/goproc && go install

ENTRYPOINT /go/bin/goproc

EXPOSE 3000
