#源镜像
FROM golang:1.17

#设置工作目录
WORKDIR $GOPATH/src

#拉取项目代码
RUN git clone https://github.com/Lastly1999/fast-admin-service

#添加配置文件
ADD ./config/app.ini $GOPATH/src/fast-admin-service

#切换工作目录
WORKDIR $GOPATH/src/fast-admin-service

#设置代理
ENV GOPROXY https://goproxy.io

#go构建可执行文件
RUN go build .

#暴露端口
EXPOSE 5600:5600

#最终运行docker的命令
ENTRYPOINT  ["./fast-admin-service"]
