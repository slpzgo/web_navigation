package handle

import (
	"github.com/labstack/echo"
	"net/http"
	"userauth/daos"
	"userauth/vm"
)



func Category(c echo.Context) error{

	data := daos.GetCategoryList()
	list := &vm.CategoryList{}
	for _,v := range data{
		c:= vm.Category{
			Id:v.Id,
			Name:v.Name,
			Description:v.Description,
		}
		list.List=append(list.List,c)
	}
	return c.JSON(http.StatusOK,list)
}


