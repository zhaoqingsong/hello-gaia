package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/log4go"
	"github.com/zhaoqingsong/studyweb/initial/config"
	"github.com/zhaoqingsong/studyweb/service"
	"github.com/zhaoqingsong/studyweb/utils"
)

type getFileAssets struct {
	ID   int
	Name string
	Url  string
}

// Resp .
type Resp struct {
	Result  int
	Message string
	Url     string
	Data    interface{}
}

//GetFileList
func GetFileList(ctx *gin.Context) {
	method := "GetFileList"

	fmt.Println("进入了controller。")
	var req getFileAssets
	err := ctx.Bind(&req)
	if err != nil {
		log4go.Info("in all caps: %q\n", "请求错误")
		return
	}
	err = service.GetFile(
		req.ID,
		req.Name,
	)

	utils.ComposerJSON(method, ctx, err, nil, nil)
	return
}

type uploadFileAssets struct {
	Name  string `form:"name"`
	Group string `form:"group"`
}

// 上传文件
func UploadFileList(ctx *gin.Context) {
	method := "UploadFileList"

	// Upload the file to specific dst.
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, config.Config.Storage.Dir+"/"+file.Filename)
	if err != nil {
		log4go.Info("[Controller] UploadFileList error in SaveUploadedFile: %s\n", err)
		return
	}
	log4go.Info("上传文件成功，filename： %s, url: %s", file.Filename, "http://download/")


	service.CreateFile(file.Filename, "http://download.hellobile.cn/download/")
	utils.ComposerJSON(method, ctx, err, nil, nil)
}
