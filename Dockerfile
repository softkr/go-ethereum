# 在goland环境下对链程序进行编译
FROM golang:1.16.9-alpine as builder
# 换源和安装依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache make gcc musl-dev linux-headers git
# 将项目录添加到容器
ADD . /go-ethereum
# 执行make编译链程序
RUN cd /go-ethereum && make aqchain && make bootnode

# 使用alpine环境作为运行
FROM alpine:latest
# 换源和安装依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache ca-certificates
# 将编译好的可执行文件复制到bin目录下
COPY --from=builder go-ethereum/build/bin/aqchain /usr/local/bin/
COPY --from=builder go-ethereum/build/bin/bootnode /usr/local/bin/
COPY entrypoint.sh /usr/local/bin/

EXPOSE 8545 8546 30301 30301/udp 30303 30303/udp

ENTRYPOINT ["entrypoint.sh"]
