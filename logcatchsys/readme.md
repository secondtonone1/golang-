## 日志采集系统
## 版本说明
tag 为 v1.0vipper 版实现了基本的日至采集功能，通过哦vipper监控配置文件修改，从而达到动态控制协程监控日志。
tag 为 v2.0etcd 版实现了基于etcd添加和删除监控日至的功能。
### golang 代理设置
如果golang下载第三方库较慢，可以使用七牛云代理
go env -w GOPROXY=https://goproxy.cn,direct
go版本要求1.13以上

### 机器需要安装zookeeper和kafka安装包
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
7 启动主程序
启动logcatchsys/logcatchsys/main.go监听文件
8 启动writefileloop循环修改文件
当循环写文件后，可以看到消费者不断打印消费日志。
9 修改config.yaml中监听的日志topic或者日志路径
当config.yaml修改后，日志采集系统动态启动协程监听新的日志路径。
