package components

import (
	"fmt"
	model "golang-/logcatchsys/web_admin/models"

	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

func InitDb() (err error) {
	user := BConfig.String("mysql::mysql_user")
	pwd := BConfig.String("mysql::mysql_pwd")
	ip := BConfig.String("mysql::mysql_ip")
	db := BConfig.String("mysql::mysql_db")
	sqlstr := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, ip, db)
	database, err := sqlx.Open("mysql", sqlstr)
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}
