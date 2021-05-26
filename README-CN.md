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
go run main.go set hello world

2021/05/25 23:13:54 Reply: OK
2021/05/25 23:13:54 Command: +OK

go run main.go get hello

2021/05/25 23:14:11 Reply: world
2021/05/25 23:14:11 Command: $5
world
```

## Docker

```shell
docker build -t="jaronnie/julenkv:v0.1" .
docker run --name=julenkv -itd -p 5200:5200 jaronnie/julenkv:v0.1
docker exec -it julenkv sh
./bin/julenkv-cli set hello world
./bin/julenkv-cli get hello 
```

## 更新记录

请在[logs.md](logs.md)上查看更多的信息。

* 2021.5.25 发布 v0.1 版本 

## License

julenkv 遵循 MIT开源协议。
