package handle

import (
	"git.bitboolean.com/platformgo/misc/globals"
	"github.com/labstack/echo"
	"net/http"
	"userauth/cache"
	"userauth/daos"
	"userauth/middleware"
	"userauth/models"
	"userauth/vm"
)

func CreateLink(c echo.Context)error{
	req := &vm.Links{}
	if err := c.Bind(req);err != nil{
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求体错误", err))
	}
	data := models.Links{
		Name:req.Name,
		Href:req.Href,
		IsVip:req.IsVip,
		Description:req.Description,
	}
	_,err := daos.CreateLink(data)
	if err !=nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("添加失败", err))
	}
	return c.JSON(http.StatusOK,req)
}

func GetUserLinks(c echo.Context)error{
	clientId := c.Request().Header.Get(middleware.UserKey)
	responseList := new(vm.LinksList)

	userList := cache.GetList(clientId,daos.GetLinks)
	for _,v := range userList{
		c := vm.Links{
			Name :v.Name,
			Href :v.Href,
			IsVip:v.IsVip,
			Description :v.Description,
			CategoryName:v.CategoryName,
		}
		responseList.List = append(responseList.List,c)
	}
	return c.JSON(http.StatusOK,responseList)

}
