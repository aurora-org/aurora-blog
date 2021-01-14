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
	tx := common.AuroraRW.Begin()

	defer func() {
		_ = tx.Close()
	}()

	if err := tx.AutoMigrate(&model.Article{}, &model.Author{}, &model.Tag{}, &model.Category{}, &model.Site{}).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Create(&model.Author{
		Name:     "wimi",
		NickName: "ass",
		Age:      22,
		Gender:   1,
		Hobby:    "sing dance and rap",
		Avatar:   "https://image.qmdx00.cn/black_avatar.jpeg",
		Extra:    "",
	}).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
