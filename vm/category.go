package vm

type Category struct {
	Id int  `json:"id"`
	Name string	 `json:"name"`
	Description string `json:"description"`
}

type CategoryList struct {
	List []Category `json:"list"`
}