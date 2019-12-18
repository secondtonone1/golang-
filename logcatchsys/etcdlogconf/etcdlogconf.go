package etcdlogconf

import (
	"context"
	kafkaqueue "golang-/logcatchsys/kafka"
)

type EtcdLogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

type EtcdLogMgr struct {
	ctx           context.Context
	cancel        context.CancelFunc
	etcdConf      *EtcdLogConf
	keyChan       chan string
	kafkaProducer *kafkaqueue.ProducerKaf
}
