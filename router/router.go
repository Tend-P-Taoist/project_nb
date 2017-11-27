package router

import (
	"github.com/gin-gonic/gin"
	"project_nb/handler"
)



func Route( engine *gin.Engine){


	//engine.Use(handler.AuthHandler)

	//静态文件路径
	//
	engine.Static("/static","/Users/ydf/gopath/src/project_nb/static")

	//模板路径
	engine.LoadHTMLGlob("/Users/ydf/gopath/src/project_nb/view/templete/**")


	router := engine.Group("/app")
	{
		router.Use(handler.AuthHandler)
		router.POST("/register", handler.RegisterHandler)
		router.POST("/registeremail", handler.RegisterByEmailHandler)
		router.GET("/activate/:uuid", handler.ActivateAccountHandler)
		router.POST("/login", handler.LoginHandler)
		router.GET("/sendMessage/:to", handler.SendCodeHandler)
		router.POST("/upload", handler.UploadHanler)

	}


	admin := engine.Group("admin")
	{
		admin.GET("portal",handler.LoadPortal)
	}
}



