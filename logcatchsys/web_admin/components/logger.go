package components

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var BConfig config.Configer = nil
var Weblogkey string

func init() {
	var err error
	BConfig, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("config init error:", err)
		return
	}

	Weblogkey = BConfig.String("log::log_key")
}

func InitLogger() (err error) {
	if BConfig == nil {
		err = errors.New("beego config new failed!")
		return
	}
	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = 1000
	}

	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::log_path")
	level, _ := BConfig.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxlines

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(confStr))
	logs.SetLogFuncCall(true)
	return
}
