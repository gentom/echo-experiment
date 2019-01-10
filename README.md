# echo-experiment
#### Recipe
```
$ cd echo-experiment/
$ go mod init
$ go get -u github.com/labstack/echo
$ vim server.go
$ go build
$ ./echo-experiment
```
#### Curl
```
$ curl http://localhost:8080/
ʕ◔ϖ◔ʔ

$ curl -X POST -H "Content-Type: application/json" -d '{"name":"jojo"}'  http://localhost:8080/post
{"name":"jojo"}
```