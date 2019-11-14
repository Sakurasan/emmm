package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Num        int    `gorm:"AUTO_INCREMENT"`
	UserId     string `gorm:"unique;not null"`
	Name       string //姓名
	Department string `gorm:"index:department"` //部门
}
