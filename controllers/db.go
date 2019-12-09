package controllers

import (
	"emmm/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func InitDb() {
	var err error
	DB, err = gorm.Open("sqlite3", "emmm.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	//defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&models.Info{}, &models.User{})
	// DB.AutoMigrate(&models.User{})
	fmt.Println("DB初始化成功")
}

func InitMysqlDb(conf map[string]string) {
	dbType := "mysql"
	dbName := conf["name"]
	user := conf["user"]
	password := conf["password"]
	host := conf["host"]

	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	DB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Fatalf("connect %s has err:%s", dbType, err)
	}
	DB.AutoMigrate(&models.Info{}, &models.User{})

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)

	fmt.Println("数据库初始化完毕", dbType)
}

//   // 创建
//   db.Create(&Product{Code: "L1212", Price: 1000})

//   // 读取
//   var product Product
//   db.First(&product, 1) // 查询id为1的product
//   db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

//   // 更新 - 更新product的price为2000
//   db.Model(&product).Update("Price", 2000)

//   // 删除 - 删除product
//   db.Delete(&product)
//   db.Delete(&product)
//   db.Delete(&product)
