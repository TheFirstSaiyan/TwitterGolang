package models

import "gorm.io/gorm"

type Follows struct {
	gorm.Model
	SourceUser string `json:"sourceuser"`
	TargetUser string `json:"targetuser"`
}
