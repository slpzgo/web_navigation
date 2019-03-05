package handle

import (
	"github.com/labstack/echo"
	"net/http"
	"userauth/cache"
	"userauth/middleware"
	"userauth/vm"
)

func SaveUserRedirect(c echo.Context) error{
	clientId := c.Request().Header.Get(middleware.UserKey)
	u := new(vm.RedirectSave)
	if err := c.Bind(u); err != nil {
		return nil
	}
	cache.Increment(clientId,u.Redirect)
	return c.JSONBlob(http.StatusOK,[]byte(`{"message":"ok"}`))

}
