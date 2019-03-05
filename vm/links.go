package vm

type Links struct {
	Name string `json:"name"`
	Href string `json:"href"`
	IsVip int `json:"is_vip"`
	Description string `json:"description"`
	CategoryName string `json:"category_name"`
}

type LinksList struct {
	List []Links `json:"list"`
}