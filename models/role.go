package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Code int
	Name string
}



