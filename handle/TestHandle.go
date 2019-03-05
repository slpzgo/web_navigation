package handle

import (
	"github.com/labstack/echo"
	"userauth/cache"
)


func TestHandle(c echo.Context) error {
	_,err := cache.RedisClient.Do("set","abc","123")
	if(err != nil){
		panic(err)
	}
	return 	c.String(200,"acb")
}
