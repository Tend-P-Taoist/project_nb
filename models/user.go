package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {

	RealName string
	NickName string
	Salt string
	Account string
	Password string
	PhoneNumber string
	Email string
	Age int
	Gender int
	BirthDay time.Time
	HeadPic string
	Active bool
	InUse bool

	gorm.Model
}

func (u User)TableName() string{
	return "user"
}