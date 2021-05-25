# English | [Chinese](README-CN.md)

## Introduction

julenkv means "just learn key-value database".

julenkv modeled on redis. Just to learn redis.

Only used for learning, never used in a production environment.

## Usage

### server

```shell
cd github.com/jaronnie/julenkv/cmd/server
go run main.go
```

### client

```shell
cd github.com/jaronnie/julenkv/cmd/client
go run main.go set hello world

2021/05/25 23:13:54 Reply: OK
2021/05/25 23:13:54 Command: +OK

go run main.go get hello

2021/05/25 23:14:11 Reply: world
2021/05/25 23:14:11 Command: $5
world
```

## logs

Read more information on [logs.md](logs.md)

* release v0.1 version on 2021.5.25

## License

julenkv is licensed under the term of the [MIT License](https://github.com/jaronnie/julenkv/blob/main/LICENSE)
