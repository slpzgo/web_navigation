package daos

import (
	"userauth/models"
)

func GetCategoryList()(category []*models.Category){
	category = make([]*models.Category,0)
	//fmt.Println(category)
	mysql.Find(&category)
	//if err != nil{
	//	panic(err)
	//}
	return category

}
