package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/henrylee2cn/faygo/ext/uuid"
	"project_nb/common/define"
)


/*上传文件相关*/
func UploadHanler(c *gin.Context){

	file,err := c.FormFile("file")
	if err !=nil{
		c.AbortWithError(500,err)
		return
	}

	// Upload the file to specific dst.
	 dst :=  "/Users/ydf/Desktop/uploadtest/" + uuid.New().String()
	 if err = c.SaveUploadedFile(file, dst);err != nil {
	 	c.AbortWithError(500,err)
	 	return
	 }
	 c.JSON(200,gin.H{"code":0,"msg":define.UploadSuccess,"data":gin.H{"path":dst}})
}
