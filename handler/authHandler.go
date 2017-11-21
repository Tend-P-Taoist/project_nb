package handler

import (
	"github.com/gin-gonic/gin"
	 "../models"
	"../db"
	"net/http"
	"../common/error"
	"../common/msg"
	"crypto/md5"

)

/*授权验证*/
func AuthHandler(c *gin.Context){

	token := c.GetHeader("token")

	userName,err := db.Redis.Do("GET",token)
	if err != nil {
		c.AbortWithError(200,err)
		return
	}

	if userName == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,gin.H{"message":error.UnAuthorizedError})
		return
	}

	c.Status(200)
}

/*注册用户(用户密码注册)*/
func RegisterHandler(c *gin.Context){

	user := models.User{}
	c.Bind(&user)
	//参数验证
	if len(user.Account) == 0 ||
		len(user.Password) == 0 {

		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":error.ParameterError})
		return
	}

	_,err := db.FindUserByAccount(user.Account)
	//用户已存在
	if err == nil{
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":error.UserExsitError})
		return
	}

	user.Password = string(md5.Sum([]byte(user.Password))[:])
	err = db.CreateUser(user)

	if err != nil{
		c.AbortWithError(500,err)
		return
	}

	c.JSON(200,gin.H{"code":0,"msg":msg.RegisterSuccess})
}


/*登录*/
func LoginHandler(c *gin.Context){

}
