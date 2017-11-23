package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {

	gorm.Model
	Token string `gorm:"-"`
	RealName string
	NickName string
	Salt string
	Account string
	Password string
	PhoneNumber string
	Email string
	Age int
	Gender int
	BirthDay *time.Time
	HeadPic string
	Active bool
	Level int
	Roles []Role `gorm:"many2many:user_roles;"`
	RoleID int


}

func (u User)TableName() string{
	return "user"
}