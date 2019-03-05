package models

import "time"

type Links struct {
	Id int
	Name string
	Href string
	IsDelete  int
	CreatedAt time.Time
	UpdatedAt time.Time
	IsVip int
	Description string
	CategoryName string
}

func (m *Links) TableName() string {
	return "links"
}