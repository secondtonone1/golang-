#!/bin/bash

echo "begin to start zookeeper"
#启动zookeeper
/usr/local/zookeeper-3.4.14/bin/zkServer.sh start &
sleep 4
echo "begin to start kafka"
#启动kafka
/usr/local/kafka_2.13-2.4.0/bin/kafka-server-start.sh  /usr/local/kafka_2.13-2.4.0/config/server.properties &
sleep 4
echo "begin to start etcd"
#启动etcd
/home/secondtonone/workspace/goProject/src/golang-/logcatchsys/config-etcd-srv/etcd_start.sh &