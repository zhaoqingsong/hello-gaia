package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhaoqingsong/studyweb/router/v1"

)

// GetRouters .
func GetRouters(r *gin.Engine) *gin.Engine {

	// {domain}/api/v1/{apps}/{service}
	rgAPIV1 := r.Group("/api/v1")
	{
		v1.Files(rgAPIV1)
	}
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "AppJenkinsGolangTest is Up")
	})
	return r
}
