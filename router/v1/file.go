package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoqingsong/studyweb/controller"
)



// File .
func Files(rg *gin.RouterGroup) {

	rgAPIV1File := rg.Group("/file")

	// assets
	rgAPIV1FileAssets := rgAPIV1File.Group("/assets")

	rgAPIV1FileAssets.GET("/list", controller.GetFileList)
	rgAPIV1FileAssets.POST("/create", controller.UploadFileList)

	rgAPIV1FileAssets.PUT("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update File",
		})
	})
	rgAPIV1FileAssets.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete File",
		})
	})

}
