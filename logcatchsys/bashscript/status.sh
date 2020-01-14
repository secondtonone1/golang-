#!/bin/bash

echo "begin to check zookeeper status"
#查看zookeeper状态
/usr/local/zookeeper-3.4.14/bin/zkServer.sh status 
sleep 3
echo "begin to check  etcd status"
#查看etcd状态
etcdctl cluster-health
sleep 3
echo "begin to check  kafka status"
#查看kafka 状态
/usr/local/kafka_2.13-2.4.0/bin/kafka-topics.sh --list --zookeeper localhost:2181 
sleep 3