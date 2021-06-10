FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    PATH="/dist/bin:${PATH}"

WORKDIR /build

COPY . .

RUN go mod tidy && go build -o bin/julenkv-server main.go && go build -o bin/julenkv-cli client/main.go

WORKDIR /dist

RUN cp -r /build/bin .

EXPOSE 5200

ENTRYPOINT ["julenkv-server"]
