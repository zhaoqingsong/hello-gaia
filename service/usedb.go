package service

import (
	"fmt"
	"github.com/zhaoqingsong/studyweb/model"
	"log"
)

//根据集群的id信息或者环境信息查询集群的详细信息,支持批量
func GetFile(id int, name string) (err error) {
	var queryDb = db
	if id > 0 {
		queryDb = queryDb.Where("ID = (?)", id).Find(&model.File{})
	}
	if name != "" && len(name) > 0 {
		queryDb = queryDb.Where("name = ?", name).Find(&model.File{})
	}

	err = queryDb.Find(&model.File{}).Error
	if err != nil {
		log.Fatal(fmt.Sprintf("【DB.LOG】dao.file.GetFileList sql execute err ,err message is %s", err.Error()))
	}
	queryDb = nil
	return err
}

//查询集群总条数
func GetFileCountService(params model.File) (count int) {
	var queryDb = db
	err = queryDb.Find(&[]model.File{}).Count(&count).Error
	if err != nil {
		log.Fatal(fmt.Sprintf("【DB.LOG】dao.cluster.GetClustersCount sql execute err ,err message is %s", err.Error()))
	}
	queryDb = nil
	return count
}

//新增文件
func CreateFile(name string, url string) (res string) {
	var queryDb = db

	log.Printf("[Service] CreateFile filename:%s, url: %s", name, url)
	file := model.File{Name: name, Url: url}
	//r := queryDb.HasTable("files")
	////r := queryDb.Find(&model.File{})
	////
	////r := queryDb.NewRecord(&file)
	//log.Printf("查询是否存在：%s", r)
	queryDb.Create(&file)
	queryDb.First(&model.File{}, 1)
	queryDb = nil
	return "新增一条数据成功 "
}

// DeleteFile
func DeleteFile(params model.File) (res string) {
	var queryDb = db
	err := queryDb.Delete(&model.File{Name:params.Name})
	if err != nil {
		log.Fatal(fmt.Sprintf("【DB.LOG】service.DeleteFile. sql execute message is %s", err))
	}

	return "删除一条数据成功 "
}
