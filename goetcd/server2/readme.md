## etcd 下载安装
windows环境下解压安装包即可，在etcd文件夹内建立三个文件，etcd做集群启动
start1.bat文件
``` bat
.\etcd.exe --name etcd01 ^
--data-dir .\data\etcd01 ^
--advertise-client-urls http://127.0.0.1:2379 ^
--listen-client-urls http://127.0.0.1:2379 ^
--listen-peer-urls http://127.0.0.1:2380 ^
--initial-advertise-peer-urls http://127.0.0.1:2380 ^
--initial-cluster-token etcd-cluster-1 ^
--initial-cluster etcd01=http://127.0.0.1:2380,etcd02=http://127.0.0.1:2381,etcd03=http://127.0.0.1:2382 ^
--initial-cluster-state new
pause
```
start2.bat文件
``` bat
.\etcd.exe --name etcd02 ^
--data-dir .\data\etcd02 ^
--advertise-client-urls http://127.0.0.1:3379 ^
--listen-client-urls http://127.0.0.1:3379 ^
--listen-peer-urls http://127.0.0.1:2381 ^
--initial-advertise-peer-urls http://127.0.0.1:2381 ^
--initial-cluster-token etcd-cluster-1 ^
--initial-cluster etcd01=http://127.0.0.1:2380,etcd02=http://127.0.0.1:2381,etcd03=http://127.0.0.1:2382 ^
--initial-cluster-state new
pause
```
start3.bat文件
``` bat
.\etcd.exe --name etcd03 ^
--data-dir .\data\etcd03 ^
--advertise-client-urls http://127.0.0.1:4379 ^
--listen-client-urls http://127.0.0.1:4379 ^
--listen-peer-urls http://127.0.0.1:2382 ^
--initial-advertise-peer-urls http://127.0.0.1:2382 ^
--initial-cluster-token etcd-cluster-1 ^
--initial-cluster etcd01=http://127.0.0.1:2380,etcd02=http://127.0.0.1:2381,etcd03=http://127.0.0.1:2382 ^
--initial-cluster-state new
pause
```
启动三个节点，并且在当前目录建立data文件夹，分别建立三个文件夹
## 测试etcd启动是否成功
etcdctl cluster-health 
输出如下
``` cmd
member 19ac17627e3e396f is healthy: got healthy result from http://127.0.0.1:4379
member bf9071f4639c75cc is healthy: got healthy result from http://127.0.0.1:2379
member e7b968b9fb1bc003 is healthy: got healthy result from http://127.0.0.1:3379
cluster is healthy
```