# English | [Chinese](README-CN.md)

## Introduction

julenkv means "just learn key-value database".

julenkv models on redis. Just to learn redis.

Only use for learning, never used in a production environment.

## Usage

### server

```shell
go run main.go
```

### client

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
docker run --name=julenkv -itd -p 5200:5200 jaronnie/julenkv:v0.3
docker exec -it julenkv sh

$ julenkv-cli

julenkv-cli> set hello world 
ok
julenkv-cli> get hello
world
```

## logs

Read more information on [logs.md](logs.md)

* release v0.1 version on 2021.5.25
* release v0.1.1 version on 2021.5.27
* release v0.2 version on 2021.5.28
* release v0.3 version on 2021.6.10

## License

julenkv is licensed under the term of the [MIT License](https://github.com/jaronnie/julenkv/blob/main/LICENSE)
