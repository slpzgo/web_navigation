package middleware

import (
	"github.com/labstack/echo"
	"userauth/helper"
)

const UserKey  = "Client-Id"

func User(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		 c.Response().Header().Set(UserKey,getClientId(c))
		 return next(c)
	}
}

func getClientId(c echo.Context)string{
	if(c.Request().Header.Get(UserKey) == ""){
		key := helper.GetRandomString()
		c.Request().Header.Set(UserKey,key)
		return  key
	}

	return c.Request().Header.Get(UserKey)
}
