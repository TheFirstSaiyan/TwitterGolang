package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserName string `json:"name"`
	Content  string `json:"content"`
}
