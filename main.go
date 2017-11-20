
package main

import (
	"github.com/gin-gonic/gin"
	)

type cc struct {
	a string
}
func main() {

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	
	engine.Run(":8888")
}
