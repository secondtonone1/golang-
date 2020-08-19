package constants
import(
	"time"
)
const (
	HEART_BEAT_INTERVAL       = time.Second * 10     //心跳间隔
	CONNECT_TIMEOUT           = time.Second * 10     //连接超时
	MAX_CONNIDLE_TIME         = time.Second * 60 * 5 //最大空闲时间
	QUERY_TIME_OUT            = time.Second * 5      //查询等待最大时间间隔
	STATIC_COUNT_REDIS_EXPIRE = time.Minute * 60     //一个小时
)

const (
	DB_DATABASES = "zack_db"
	DB_COLLECTION = "zack_collection"
)