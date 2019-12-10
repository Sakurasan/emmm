package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId           string `gorm:"unique;index:uid"`
	Department       string //`gorm:"index:department"` //部门
	Name             string `gorm:"index:uname"` //姓名
	ComputerType     string //电脑型号
	OsType           string //系统类别
	MacAddr          string //mac地址
	SecuritySoftWare string //杀软名称
}

func (t User) TableName() string {
	return "user"
}
