package daos

import (
	"userauth/models"
)

func CreateLink(data models.Links)(int64,error){
	id,err := mysql.Insert(data)
	return id,err
}

func GetLinks()[]models.Links{
	list := make([]models.Links,0)
	mysql.Find(&list)
	return list

}