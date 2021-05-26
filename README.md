# English | [Chinese](README-CN.md)

## Introduction

julenkv means "just learn key-value database".

julenkv modeled on redis. Just to learn redis.

Only used for learning, never used in a production environment.

## Usage

### server

```shell
go run main.go
```

### client

```shell
cd client
go run main.go set hello world

2021/05/25 23:13:54 Reply: OK
2021/05/25 23:13:54 Command: +OK

go run main.go get hello

2021/05/25 23:14:11 Reply: world
2021/05/25 23:14:11 Command: $5
world
```

### golang

```go
package main

import (
	"fmt"
	julenkv "github.com/jaronnie/julenkv/client/go-julenkv"
)

func main() {
	conn, _ := julenkv.Connect("0.0.0.0:5200")
	reply, _ := julenkv.Do(conn, "get", "hello")
	if _, ok := reply.(string); ok {
		value := reply.(string)
		fmt.Println(value)
	}
}

```

## Docker

```shell
docker build -t="jaronnie/julenkv:v0.1.1" .
docker run --name=julenkv -itd -p 5200:5200 jaronnie/julenkv:v0.1.1
docker exec -it julenkv sh
./bin/julenkv-cli set hello world
./bin/julenkv-cli get hello 
```

## logs

Read more information on [logs.md](logs.md)

* release v0.1 version on 2021.5.25
* release v0.1.1 version on 2021.5.27

## License

julenkv is licensed under the term of the [MIT License](https://github.com/jaronnie/julenkv/blob/main/LICENSE)
