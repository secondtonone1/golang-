package config

import "sync"

const (
	STATUS_SEC_SUCCESS       = 1000
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
	STATUS_PRODUCTID_INVALID = 1011 //商品id参数不对
	STRING_CONVERT_FAILED    = 1012 //字符串转换失败
	AUTH_SIGN_CHECK_FAILED   = 1013 //user auth sign check failed
	USER_ID_INVALID          = 1014 //user id 错误
	FREQUENCY_LIMIT          = 1015 //每秒限流，用户抢购次数过多
	IP_LIMIT                 = 1016 //每秒同一个ip抢购太多
	MSG_CHAN_CLOSED          = 1017 //msg chan 关闭
	STATUS_REQ_TIMEOUT       = 1018 //超时
	FROM_REDIS_CHAN_CLOSED   = 1019 //from redis chan 关闭
)

type SecKillConf struct {
	EtcdConfData      EtcdConf
	RedisBlacklist    RedisConf
	LogConfData       LogConf
	SecInfoData       map[int]*SecInfoConf
	SecInfoRWLock     sync.RWMutex    //secinfodata的读写锁
	CookieSecretKey   string          //抢购认证秘钥
	FrequencyLimit    int             //用户访问每秒频率限制
	IpLimit           int             //ip访问每秒频率限制
	ReferWhitelist    []string        //refer跳转白名单
	IDBlacklist       map[int]bool    //usr id 黑名单
	IPBlacklist       map[string]bool // ip 黑名单
	BlacklistRWLock   sync.RWMutex    //黑名单的读写锁
	RedisReadGoCount  int             //读取redis goroutine 数量
	RedisWriteGoCount int             //写入redis goroutine 数量
}

type SecInfoConf struct {
	ProductId int   `json:"productid" db:"productid"`
	StartTime int64 `json:"starttime" db:"starttime"`
	EndTime   int64 `json:"endtime" db:"endtime"`
	Status    int   `json:"status" db:"status"`
	Total     int   `json:"total" db:"total"`
	Left      int   `json:"_" db:"left"`
}

type SecRequest struct {
	ProductId    int
	Source       string //来源
	AuthCode     string //校验码
	SecTime      string //抢购时间
	Nance        string //随机数
	UserId       int    // 用户id
	UserAuthSign string //用户cookie
	SecTimeStamp int64  // 访问时间戳
	ClientAddr   string //客户端ip
	ReferAddr    string //哪个地址跳转的
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

type RedisMgrConf struct {
	ReadFromRedisGrtCount int
	WrtieToRedisGrtCount  int
}
