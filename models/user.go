package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName" gorm:"user_name" form:"userName" validate:"required,gt=2"`
	Password string `json:"password" gorm:"password" form:"password" validate:"required,gt=5,lt=32"`
	Ip       string `json:"ip" gorm:"ip"`
}
