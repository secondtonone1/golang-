package config

import "sync"

const (
	STATUS_SELL_NORMAL       = 1000
	STATUS_SELL_FORBID       = 1001 //禁止销售
	STATUS_SELL_OUT          = 1002 //售罄
	STATUS_SELL_NOT_BEGIN    = 1003 //活动未开始
	STATUS_SELL_END          = 1004 //活动已结束
	STATUS_SELL_NOT_ENOUGH   = 1005 //数量不足
	STATUS_PRODUCT_NOT_FOUND = 1006 //商品未找到
	STATUS_PRODUCT_TIME_ERR  = 1007 //商品时间错误
	JSON_MARSHAL_ERR         = 1008 //JSON序列化失败
	CONVERT_PRODUCT_INFO_ERR = 1009 //PRODUCT 信息生成失败
	STATUS_PRODUCT_LIST_ERR  = 1010 //product list 获取失败
)

type SecKillConf struct {
	EtcdConfData EtcdConf
	RdisConfData RedisConf
	LogConfData  LogConf
	SecInfoData  map[int]*SecInfoConf
	SecInfoRLock sync.RWMutex
}

type SecInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}

type RedisConf struct {
	RedisAddr      string
	RedisMaxIdle   int
	RedisMaxActive int
	RedisIdleTime  int
}

type EtcdConf struct {
	EtcdAddr       string
	EtcdTimeout    int
	EtcdSecPrefix  string
	EtcdSecProduct string
}

type LogConf struct {
	LogLv    int    `json:"level"`
	LogPath  string `json:"filename"`
	MaxLines int    `json:"maxlines"`
}
