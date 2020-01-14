#!/bin/bash
ID=`ps -ef | grep "etcd" | grep -v "grep" | grep -v "etcd_stop"| awk '{print $2}'`
echo $ID
echo "---------------"
for id in $ID
do
kill -9 $id
echo "killed $id"
sleep 2
done
echo "---------------"
sleep 3
echo "begin to delete etcd data"
rm -rf /home/secondtonone/workspace/goProject/src/golang-/logcatchsys/config-etcd-srv/data/


#ps -ef|grep etcd|grep -v grep|awk '{print $2}'|xargs kill -9
#sleep 3
#rm -rf ./data