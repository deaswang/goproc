# goproc

goproc is an RESTful api server for view linux /proc path file information.

goproc 是读取 linux proc 信息的 RESTful API。

## Install

```bash
go get github.com/deaswang/goproc
cd $GOPATH/src/github.com/deaswang/goproc
go install
goproc
```

```bash
Usage of ./goproc:
  -cert string
        The cert file name for tls. (default "ssl.csr")
  -key string
        The key file name for tls. (default "ssl.key")
  -port int
        The proxy server port. (default 3000)
  -token string
        The token file for authentication. (default "token.txt")
```

add token to token.txt for auth client. use postman to open doc/goproc.postman_collection.json test

添加认证 token 到 token.txt。使用 postman 打开 doc/goproc.postman_collection.json 测试。
