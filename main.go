package main

import (
	"flag"
	"fmt"
	"git.bitboolean.com/platformgo/misc/globals"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"userauth/boostrap"
	"userauth/config"
	"userauth/handle"
	selfMiddleware "userauth/middleware"
)

var (
	cfg = flag.String("c", "cfg.yml", "The path of configuration file")
)

func init() {
	flag.Parse()
	fmt.Println(cfg)
	bootstrap.Bootstrap(*cfg)
}

func main() {
	defer bootstrap.Destroy()

	e := echo.New()

	e.Static("/static", "static")
	e.File("/", "static/Index.html")
	e.File("/admin", "static/admin.html")

	e.Validator = globals.NewDefaultValidator()
	route(e)
	e.Start(config.Config().TestUrl+config.Config().HTTPBind)
}



func route(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	//api
	group := e.Group("/api/v1/",selfMiddleware.User)
	group.GET("test",handle.TestHandle)
	group.GET("links",handle.GetUserLinks)
	group.POST("redirect",handle.SaveUserRedirect)


	////admin
	adminGroup := e.Group("/api/v1/admin/",selfMiddleware.Auth)
	adminGroup.GET("category",handle.Category)
	adminGroup.POST("link",handle.CreateLink)

}
