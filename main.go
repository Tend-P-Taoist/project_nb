
package main

import (
	"github.com/gin-gonic/gin"
	)

func main() {

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	engine.Run(":8888")
}
