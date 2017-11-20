package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

var DB gorm.DB

func init()  {
	db,err := gorm.Open("mysql","sifan:sifan@/project_nb?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
}
