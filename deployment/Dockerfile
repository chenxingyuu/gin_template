# 指定基础镜像和编译平台
FROM --platform=linux/amd64 golang:1.21-alpine AS builder
# 换源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 下载依赖
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /builder
WORKDIR /builder
ENV GOPROXY=https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
# 设置 GOOS 和 GOARCH 环境变量
ENV GOOS=linux
ENV GOARCH=amd64
COPY . .
RUN go build -o ./server /builder/cmd/server/main.go

# 创建最终镜像
FROM alpine:latest
# 换源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 下载依赖
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*
# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /
COPY --from=builder /builder/server .
CMD ["./server"]
