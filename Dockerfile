FROM golang:alpine
MAINTAINER RongXinChen "1358826554@qq.com"
WORKDIR /fast-admin-service
ADD ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 指定编译完成后的文件名，可以不设置使用默认的，最后一步要执行该文件名
RUN go build -o fast-service .
EXPOSE 8080
# 这里跟编译完的文件名一致
ENTRYPOINT  ["./fast-service"]