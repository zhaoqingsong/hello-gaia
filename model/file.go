package model

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Name string
	Url string
}


//// 将 File 的表名设置为 `file`
//func (File) TableName() string {
//	return "File"
//}

