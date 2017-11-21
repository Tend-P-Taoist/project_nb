
package main

import (
	"github.com/gin-gonic/gin"
	_ "./db"
	"./router"

)

func main() {

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	//配置路由
	router.Route(engine)

	engine.Run(":8888")
}
