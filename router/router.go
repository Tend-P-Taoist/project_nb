package router

import (
	"github.com/gin-gonic/gin"
	"../handler"
)



func Route( engine *gin.Engine){


	//engine.Use(handler.AuthHandler)

	//静态文件路径
	engine.Static("/static","../static")

	//模板路径
	//engine.LoadHTMLGlob("../view/templete/**")

	engine.POST("/register",handler.RegisterHandler)
}
