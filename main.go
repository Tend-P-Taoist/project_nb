
package main

import (
	"github.com/gin-gonic/gin"
	_ "project_nb/db"
	"project_nb/router"

)

func main() {

	//gin.SetMode(gin.DebugMode)


	//设置日志
	//f,err := os.Create("gin.log")
	//if err != nil{
	//	panic(err)
	//}

	//gin.DefaultWriter = io.MultiWriter(f)

	engine := gin.Default()

	//配置路由
	router.Route(engine)

	engine.Run(":8888")
}
