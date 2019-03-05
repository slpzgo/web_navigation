package bootstrap

import (
	"git.bitboolean.com/platformgo/misc/driver"
	"github.com/garyburd/redigo/redis"
	"userauth/cache"
	"userauth/config"
	"userauth/daos"
)

//Bootstrap 配置参数，启动相关组件
func Bootstrap(cfg string) {
	parseConfig(cfg)
	initMysql()
	initRedis()
}

func parseConfig(cfg string) {
	config.ParseConfig(cfg)
}

func initMysql() {
	engine, err := driver.CreateMysql(config.Config().Mysql)
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(20)
	engine.SetMaxOpenConns(10)
	daos.SetMySql(engine)
}

func Destroy() {
	daos.CloesMySQL()
}

func initRedis(){
	RedisClient , err  := redis.Dial("tcp",config.Config().Redis.Host+ config.Config().Redis.Port)
	if err != nil {
		panic(err)
	}
	cache.SetInit(RedisClient)

}

