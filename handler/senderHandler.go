package handler

import (
	"github.com/gin-gonic/gin"
	"../common/define"
	"../common/sender"
	"../db"
	"strings"
	"regexp"
	"math/rand"
)


func SendCodeHandler(c *gin.Context){

	to := strings.TrimSpace(c.Param("to"))

	if len(strings.TrimSpace(to)) == 0 {
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.ParameterError})
	}
	matchPhone,err := regexp.MatchString(define.PhonePattern,to)
	if err != nil{
		c.AbortWithError(500,err)
		return
	}
	matchEmail,err := regexp.MatchString(define.EmailPattern,to)
	if err != nil{
		c.AbortWithError(500,err)
		return
	}

	if !(matchEmail || matchPhone){
		c.AbortWithStatusJSON(200,gin.H{"code":1,"msg":define.ParameterError})
		return
	}

	code := rand.Intn(10000)
	if matchEmail {
		go sender.SendEmail(to,string(code))
	}
	if matchPhone {
		go sender.SendMessage(to,string(code))
	}

	db.Redis.Do("SET",to,code)

	c.JSON(200,gin.H{"code":0,"msg":define.MessageSendSuccess})
}