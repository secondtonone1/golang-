FROM golang:latest
 
MAINTAINER zack  <secondtonone1@163.com>

RUN mkdir -p /home/secondtonone/workspace/dockerdemo
COPY ./  /home/secondtonone/workspace/dockerdemo
WORKDIR /home/secondtonone/workspace/dockerdemo
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver.bin webserver.go
CMD  /home/secondtonone/workspace/dockerdemo/webserver.bin
