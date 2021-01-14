package main

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 初始化数据库，创建对应数据表
func main() {
	common.SetupConfig()
	common.SetupDB()
	db := common.AuroraRW

	defer func() {
		_ = db.Close()
	}()

	if err := db.AutoMigrate(&model.Article{}, &model.Author{}, &model.Category{}, &model.Site{}, &model.Tag{}).Error; err != nil {
		log.Fatal(err)
	}
}
