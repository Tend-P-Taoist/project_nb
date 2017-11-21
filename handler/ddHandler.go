package handler

import (
	"github.com/gin-gonic/gin"
	"../models"
	"fmt"
	"../db"
)


func DHandler(c *gin.Context){
	users := []models.User{}
	db.DB.Where("age=?",14).Find(&users)
	num := len(users)
	fmt.Print(num)

	c.HTML(200,"hello.tmpl",gin.H{"hello":"hello_tmpl"})
}