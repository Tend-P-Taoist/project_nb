package handler

import (
	"github.com/gin-gonic/gin"
	 "project_nb/models"
	"project_nb/db"
	"net/http"
	"project_nb/common/define"
	"project_nb/common/sender"
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
	"github.com/henrylee2cn/faygo/ext/uuid"
	"github.com/garyburd/redigo/redis"

)

/*授权验证*/
func AuthHandler(c *gin.Context){

	token := c.GetHeader("token")

	user,err := db.Redis.Do("GET",token)
	if err != nil {
		c.AbortWithError(500,err)
		return
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,gin.H{"code":403,"msg":define.UnAuthorizedError})
		return
	}

	c.Next()
}

/*注册用户(用户密码注册)*/
func RegisterHandler(c *gin.Context){

	user := models.User{}
	c.Bind(&user)
	//参数验证
	if len(user.Account) == 0 ||
		len(user.Password) == 0 {

		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.ParameterError})
		return
	}

	users,err := db.FindUserByAccount(user.Account)

	if err != nil {
		c.AbortWithError(500,err)
	}
	//用户已存在
	if len(users) > 0{
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.UserExistError})
		return
	}

	hash := md5.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
	err = db.CreateUser(&user)

	if err != nil{
		c.AbortWithError(500,err)
		return
	}

	c.JSON(200,gin.H{"code":0,"msg":define.RegisterSuccess})
}

/*注册用户(邮箱注册)*/
func RegisterByEmailHandler(c *gin.Context){

	user := models.User{}
	c.Bind(&user)

	//参数验证
	match,err := regexp.MatchString(define.EmailPattern,strings.TrimSpace(user.Email))
	if err != nil {
		c.AbortWithError(200,err)
		return
	}

	if !match {
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.InvalidEmailError})
		return
	}
	if len(user.Password) == 0 {

		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.ParameterError})
		return
	}


	users,err := db.FindUserByEmail(user.Email)

	if err != nil {
		c.AbortWithError(500,err)
	}
	//用户已存在
	if len(users) > 0 {
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.UserExistError})
		return
	}

	//md5密码
	hash := md5.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
	//user.BirthDay = time.Now()

	//需要激活
	//user.Active = false
	err = db.CreateUser(&user)

	if err != nil{
		c.AbortWithError(500,err)
		return
	}

	uuid := uuid.New().String()

	_,err = db.Redis.Do("SET",uuid,user.ID)
	if err != nil {
		c.AbortWithError(500,err)
		return
	}


	//go sender.SendEmail(user.Email,"127.0.0.1:8888/activate/" + uuid)
	go sender.SendEmail(user.Email,`
    this is link <a href="127.0.0.1:8888/activate/`+ uuid +`">127.0.0.1:8888/activate/`+ uuid +`</a>
`)
	c.JSON(200,gin.H{"code":0,"msg":define.RegistEmailSuccess})
}

/*激活邮箱*/
func ActivateAccountHandler(c *gin.Context){

	uuid := c.Param("uuid")
	if len(uuid) == 0{
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.ParameterError})
		return
	}


	id,err := redis.Int(db.Redis.Do("GET",uuid))
	if err != nil {
		c.AbortWithError(500,err)
		return
	}

	//激活账号
	err = db.ActivateUserById(id)
	if err != nil {
		c.AbortWithError(500,err)
		return
	}
	db.Redis.Do("DEL",uuid)

	c.JSON(200,gin.H{"code":200,"msg":define.ActivateSuccess})
}


/*登录(密码登录)*/
func LoginHandler(c *gin.Context){

	user := models.User{}
	c.Bind(&user)
	code := c.Param("code")
	account := strings.TrimSpace(user.Account)
	var users []models.User

	//邮箱登录
	match,err := regexp.MatchString(define.EmailPattern,account)
	if err != nil {
		c.AbortWithError(200,err)
		return
	}

	if match {
		c.Bind(&user)
		users,err = db.FindUserByEmail(account)
		if err != nil {
			c.AbortWithError(500,err)
			return
		}

	}

	//手机号登录
	match,err = regexp.MatchString(define.PhonePattern,account)
	if err != nil {
		c.AbortWithError(200,err)
		return
	}

	if match {
		c.Bind(&user)
		users,err = db.FindUserByPhoneNumber(account)
		if err != nil {
			c.AbortWithError(500,err)
			return
		}

	}

	if len(users) == 0 {
		users,err = db.FindUserByAccount(account)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}
	//用户不存在
	if len(users) == 0 {
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.UserNotExistError})
		return
	}

	//密码登录
	if len(code) == 0{
		//md5密码
		hash := md5.New()
		hash.Write([]byte(user.Password))
		pwd := hex.EncodeToString(hash.Sum(nil))


		//密码错误
		if users[0].Password != pwd{
			c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.WrongPasswordError})
			return
		}
	}

	//验证码登录
	lcode,err := redis.String(db.Redis.Do("GET",account))
	if err != nil {
		c.AbortWithError(500,err)
		return
	}
	if code != lcode{
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.InvalidCodeError})
		return
	}
	//未激活
	if !users[0].Active{
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.InactiveError})
		return
	}

	//删除验证码
	db.Redis.Do("DEL",account)

	//登录成功
	c.JSON(200,gin.H{"code":0,"msg":define.LoginSuccess,"data":users[0]})
	return

}


/*个人信息修改*/
func ModifyUserHanler(c *gin.Context){

	user:= models.User{}
	c.Bind(&user)

	if err := db.ModifyUser(&user); err != nil{
		c.AbortWithError(500,err)
		return
	}


}


/*权限更改*/
func ChangeAuthorityHandler(c *gin.Context){

}
