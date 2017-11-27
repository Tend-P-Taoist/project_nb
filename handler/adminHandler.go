package handler

import "github.com/gin-gonic/gin"

func LoadPortal(c *gin.Context){

	c.HTML(200,"portal.html",[]string{"AA","BB"})
}