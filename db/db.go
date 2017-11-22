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
func FindUserByAccount(account string) ([]models.User,error){

	users := []models.User{}
	c := DB.Find(&users,"account = ?",account)
	if c.Error != nil{
		return nil,c.Error
	}
	return users,nil
}

/*根据EMAIL查找用户*/
func FindUserByEmail(email string) ([]models.User,error){

	users := []models.User{}
	c := DB.Find(&users,"email = ?",email)

	if c.Error != nil{
		return nil,c.Error
	}

	return users,nil
}

/*根据手机查找用户*/
func FindUserByPhoneNumber(num string) ([]models.User,error){

	users := []models.User{}
	c := DB.Find(&users,"phone_number = ?",num)

	if c.Error != nil{
		return nil,c.Error
	}

	return users,nil
}




func CreateUser(user *models.User) error{

	c := DB.Create(user)

	return  c.Error
}

func ActivateUserById(id int) error{

	c := DB.Model(&models.User{}).Where("id = ?",id).Update("active",true)

	return c.Error
}