FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o bin/julenkv-server cmd/server/main.go

RUN go build -o bin/julenkv-cli cmd/client/main.go

WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp -r /build/bin .

EXPOSE 5200

CMD ["/dist/bin/julenkv-server"]
