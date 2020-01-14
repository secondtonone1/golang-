## 日志采集系统
## 版本说明
tag 为 v1.0vipper 版实现了基本的日至采集功能，通过哦vipper监控配置文件修改，从而达到动态控制协程监控日志。
tag 为 v2.0etcd 版实现了基于etcd添加和删除监控日至的功能。
### golang 代理设置
如果golang下载第三方库较慢，可以使用七牛云代理
go env -w GOPROXY=https://goproxy.cn,direct
go版本要求1.13以上

### windows环境 机器需要安装zookeeper和kafka安装包
1 设置zookeeper配置文件，conf文件夹下新建zoo.cfg
``` cfg
# The number of milliseconds of each tick
tickTime=2000
# The number of ticks that the initial 
# synchronization phase can take
initLimit=10
# The number of ticks that can pass between 
# sending a request and getting an acknowledgement
syncLimit=5
# the port at which the clients will connect
clientPort=2181
dataDir=D:\\kafkazookeeper\\zookeeper-3.4.14\\data
dataLogDir=D:\\kafkazookeeper\\zookeeper-3.4.14\\log
```
设置dataDir(数据目录),dataLogDir(日志目录),clientPort(客户端端口)
2 进入zookeeper文件夹，点击bin目录下的zkServer.cmd启动zookeeper
3 设置kafka配置，修改config目录下server.properties文件，添加
log.dirs=D:\\kafkazookeeper\\kafka_2.12-2.2.0\\logs

### 启动和测试
4 启动kafka，执行如下命令
.\bin\windows\kafka-server-start.bat .\config\server.properties
5 创建测试的topic
``` cmd
.\bin\windows\kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 16 --topic logdir1
.\bin\windows\kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 16 --topic logdir2
.\bin\windows\kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 16 --topic logdir3
```
6 启动消费者
``` cmd
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic logdir1 --from-beginning
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic logdir2 --from-beginning
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic logdir3 --from-beginning
```
7 启动etcd
记得设置etcdctl版本为3  set ETCDCTL_API=3

8 启动主程序
启动logcatchsys/logcatchsys/main.go监听文件

9 启动writefileloop循环修改文件
当循环写文件后，可以看到消费者不断打印消费日志。

10 修改config.yaml中监听的日志topic或者日志路径
当config.yaml修改后，日志采集系统动态启动协程监听新的日志路径。

11 通过etcdwrite文件夹下etcdwrite.go可以更新etcd中键值为collectlogkey1的数据，从而观察我们的日志系统根据日志路径修改，动态启动和关闭协程监控。


### 补充下 Linux 环境安装Zookeeper
同样是将压缩包解压至usr/local目录下
tar zxvf zookeeper-3.4.14.tar.gz -C /usr/local
然后复制配置文件/usr/local/zookeeper-3.4.14/conf/zoo_sample.cfg为/usr/local/zookeeper-3.4.14/conf/zoo.cfg
接下来编辑/usr/local/zookeeper-3.4.14/conf/zoo.cfg
``` cmd
dataDir=/home/secondtonone/zookeeper/data
dataLogDir=/home/secondtonone/zookeeper/log
```
启动zookeeper
/usr/local/zookeeper-3.4.14/bin/zkServer.sh start
查看zookeeper 状态
/usr/local/zookeeper-3.4.14/bin/zkServer.sh status
停止zookeeper 
/usr/local/zookeeper-3.4.14/bin/zkServer.sh stop
### 补充下 Linux环境下安装kafka
下载kafka，然后解压
tar zxvf ./kafka_2.13-2.4.0.tgz  -C /usr/local/
可以看到属主是root，我将它属主变为我自己的账户
chown -R secondtonone ./kafka_2.13-2.4.0
然后我们修改kafka配置
``` cmd
broker.id=0 
port=9092 #端口号 
host.name=localhost #单机可直接用localhost
log.dirs=/home/secondtonone/kafka/logs #日志存放路径可修改可不修改
zookeeper.connect=localhost:2181 #zookeeper地址和端口，单机配置部署，localhost:2181 
```
启动kafka
/usr/local/kafka_2.13-2.4.0/bin/kafka-server-start.sh  /usr/local/kafka_2.13-2.4.0/config/server.properties

### 配置elastic和启动


