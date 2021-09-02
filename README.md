# goproc

[![Go Report Card](https://goreportcard.com/badge/github.com/deaswang/goproc)](https://goreportcard.com/report/github.com/deaswang/goproc)

goproc is an RESTful api server for view linux /proc path file information.

goproc 是读取 linux proc 路径内文件信息的 RESTful API。

## Install

```bash
go get github.com/deaswang/goproc
cd $GOPATH/src/github.com/deaswang/goproc
go install
goproc
```

## Usage

```
Usage of goproc:
  -cert string
        The cert file name for tls. (default "ssl.csr")
  -key string
        The key file name for tls. (default "ssl.key")
  -port int
        The proxy server port. (default 8809)
  -token string
        The token file for authentication. (default "token.txt")
```

add token to token.txt for auth client. if token.txt file not exist, no auth for request.
auth use request Header field token.

添加认证 token 到 token.txt。如果没有 token.txt 文件，请求不需验证。验证使用请求 token Header 值。

use postman to import doc/goproc.postman_collection.json and goproc.postman_environment.json test.

使用 postman 导入 doc/goproc.postman_collection.json 和 goproc.postman_environment.json 测试。

## API

```
GET /cpuinfo
GET /buddyinfo
GET /diskstats
GET /interrupts
GET /loadavg
GET /locks
GET /meminfo
GET /misc
GET /mounts
GET /partitions
GET /snmp
GET /softirqs
GET /stat
GET /uptime
GET /version
GET /vmstat
GET /processes
GET /net

GET /{pid}  for example: GET /666
```

It will return json response or error.

## Docker

Build docker at root path.

```bash
docker build -t goproc .
```

Run docker container.

```bash
docker run -d -p 8809:8809 goproc
```

Use docker may lack of some proc file info.
