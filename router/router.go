package router

import (
	"github.com/gin-gonic/gin"
	"project_nb/handler"
)



func Route( engine *gin.Engine){


	//engine.Use(handler.AuthHandler)

	//静态文件路径
	engine.Static("/static","../static")

	//模板路径
	//engine.LoadHTMLGlob("../view/templete/**")


	router := engine.Group("/app")

		router.POST("/register",handler.RegisterHandler)
		router.POST("/registeremail",handler.RegisterByEmailHandler)
		router.GET("/activate/:uuid",handler.ActivateAccountHandler)
		router.POST("/login",handler.LoginHandler)
		router.GET("/sendMessage/:to",handler.SendCodeHandler)
		router.POST("/upload",handler.UploadHanler)

}

