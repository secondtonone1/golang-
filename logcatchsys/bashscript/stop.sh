#!/bin/bash
# 关闭etcd

echo "begin to stop etcd"
/home/secondtonone/workspace/goProject/src/golang-/logcatchsys/config-etcd-srv/etcd_stop.sh &
sleep 3
# 关闭kafka
echo "begin to stop kafka"
/usr/local/kafka_2.13-2.4.0/bin/kafka-server-stop.sh &
sleep 3
# 关闭zookeeper
echo "begin to stop zookeeper"
/usr/local/zookeeper-3.4.14/bin/zkServer.sh stop &
sleep 3
