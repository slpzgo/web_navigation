package daos

import (
	"github.com/go-xorm/xorm"
)

var mysql *xorm.Engine


// SetDriver 设置驱动
func SetMySql(e *xorm.Engine) {
	mysql = e

}

func CloesMySQL() {
	mysql.Close()
}
