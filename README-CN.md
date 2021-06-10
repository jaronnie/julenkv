## 简介

julenkv意思是 “just learn key-value database”。

julenkv模仿redis的实现，只是为了学习 redis。

仅用于学习，请不要在生产环境中使用。

## 使用

### server 端

```shell
go run main.go
```

### client 端

```shell
cd client
go run main.go

julenkv-cli> set hello world
ok
julenkv-cli> get hello
world
```

### golang

```go
package main

import (
	"fmt"
	"github.com/jaronnie/julenkv/client/go-julenkv"
)

func main() {
	conn, _ := julenkv.Connect("0.0.0.0:5200")
	result, _ := conn.Do("set", "nj", "jay")
	fmt.Println(result)
	result, _ = conn.Do("get", "nj")
	fmt.Println(result)
}
```

## Docker

```shell
docker build -t="jaronnie/julenkv:v0.3" .
docker run --name=julenkv -itd -p 5200:5200 jaronnie/julenkv:v0.2
docker exec -it julenkv sh
./bin/julenkv-cli

julenkv-cli> set hello world 
ok
julenkv-cli> get hello
world
```

## 更新记录

请在[logs.md](logs.md)上查看更多的信息。

* 2021.5.25 发布 v0.1 版本 
* 2021.5.27 发布 v0.1.1 版本
* 2021.5.28 发布 v0.2 版本
* 2021.6.10 发布 v0.3 版本

## License

julenkv 遵循 MIT开源协议。
