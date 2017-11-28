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
	PhoneNumber string //`gorm:"unique_index";`
	Email string
	Age int
	Gender int
	BirthDay *time.Time
	HeadPic string
	Active bool
	Level int

	WxOpenid string
	WbOpenid string
	QqOpenid string

	//Roles []Role `gorm:"many2many:user_roles;"`

	RoleID int
	Role Role



}

func (u User)TableName() string{
	return "user"
}

