package user

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Age  int    `json:"id" gorm:"not null;default:0"`
}

func (*Users) TableName() string { return "users" }
