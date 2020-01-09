## 安装elasticsearch
这里简单说下Linux环境下如何安装
### 安装jdk
由于elastic是通过java实现的，所以先安装jdk，如果输入java -version显示版本号，则不用安装
我这里下载了jdk1.8版本，然后放在指定的目录解压
tar zxvf ./jdk-8u101-linux-x64.tar.gz
### 将jdk配置环境变量中
这里配置在系统环境变量中
sudo vim /etc/profile 打开系统文件
填写如下内容，然后保存
``` cmd
JAVA_HOME=/home/secondtonone/jdk/jdk1.8.0_101
JRE_HOME=$JAVA_HOME/jre
JAVA_BIN=$JAVA_HOME/bin
CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib
PATH=$PATH:$JAVA_HOME/bin:$JRE_HOME/bin
export JAVA_HOME JRE_HOME PATH CLASSPATH
```
保存后，source /etc/profile 生效
### 下载elastic并解压
``` cmd
wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.2.4.tar.gz
tar -zxvf elasticsearch-6.2.4.tar.gz -C /usr/local/
```
这时如果执行
``` cmd
sh /usr/local/elasticsearch-6.2.4/bin/elasticsearch
```
执行上述命令回报错，因为elasticsearch增加了安全设置，不能通过root执行，所以我们创建一个账号
并将该目录赋予给这个账号
``` cmd
groupadd secondtonone
useradd secondtonone -g secondtonone
chown -R secondtonone:secondtonone elasticsearch-6.2.4  
```
创建日志目录和数据目录
``` cmd
mkdir /home/secondtonone/eleticsearch/log
mkdir /home/secondtonone/eleticsearch/data
```
### 修改es配置
在/usr/local/elasticsearch-6.2.4/config/elasticsearch.yml设置数据目录，日志目录，地址和端口
``` cmd
path.data: /home/secondtonone/eleticsearch/data
path.logs: /home/secondtonone/eleticsearch/log
network.host: 127.0.0.1
http.port: 9200
```
之后/usr/local/elasticsearch-6.2.4/bin/elasticsearch 启动
## 代码测试增删改查
测试增删改查的代码在goelesearch.go中