package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"../models"
)

var DB *gorm.DB

func init()  {

	db,err := gorm.Open("mysql","sifan:sifan@/project_nb?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	//初始化数据库
	if !db.HasTable(&models.User{}){
		db.CreateTable(models.User{})
	}

	DB = db
	if err != nil {
		panic(err.Error())
	}
}


//查找用户
func FindUser(user models.User) (*models.User,error) {

	c := DB.First(&user)
	if c.Error != nil{
		return nil,c.Error
	}
	return &user,nil
}

/*根据用户名查找用户*/
func FindUserByAccount(account string) (*models.User,error){

	user := models.User{}
	c := DB.First(&user,"account = ?",account)
	if c.Error != nil{
		return nil,c.Error
	}
	return &user,nil
}


func CreateUser(user models.User) error{

	c := DB.Create(&user)
	if c.Error != nil{
		return  c.Error
	}
	return  nil
}