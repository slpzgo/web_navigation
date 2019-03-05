package middleware

import (
	"github.com/labstack/echo"
)


func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//todo
		return next(c)
	}
}



