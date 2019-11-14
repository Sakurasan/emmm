package models

import (
	"github.com/jinzhu/gorm"
)

type Info struct {
	gorm.Model
	Department       string //部门
	Name             string //姓名
	ComputerType     string //电脑型号
	OsType           string //系统类别
	MacAddr          string //mac地址
	SecuritySoftWare string //杀软名称
	ScanTime         string `gorm:"index:scantime"` //扫描时间
	ScanResult       string //扫描结果
	VirusNum         string //病毒数
	AbnormalNum      string //异常数
	HowDealWith      string //处理方式
	DealName         string
	Adviser          string
}
