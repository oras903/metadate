# 得到最新的 golang docker 镜像
FROM golang:latest
# 在容器内部创建一个目录来存储我们的应用，接着使它成为工作目录。
RUN mkdir -p /go/src/haproxy_reload
ADD . /go/src/haproxy_reload/
WORKDIR /go/src/haproxy_reload
# 下载并安装第三方依赖到容器中
RUN go get github.com/r3labs/sse
RUN go get github.com/buger/jsonparser
#编译
RUN go build -o main .
# 告诉 Docker 启动容器运行的命令
CMD ["/go/src/haproxy_reload/main"]
