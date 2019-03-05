package vm

import "time"

type user struct {
	Name  string
	ClientId string
	LastUse time.Time

}

type RedirectSave struct {
	Redirect  string `json:"category" form:"category" query:"category"`
}
