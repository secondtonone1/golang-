ps -ef|grep etcd|grep -v grep|awk '{print $2}'|xargs kill -9
sleep 3
rm -rf ./data