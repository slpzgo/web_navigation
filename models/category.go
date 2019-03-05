package models

import "time"

type Category struct {
	Id int
	Name string
	Description string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}


func (m *Category) TableName() string {
	return "category"
}


