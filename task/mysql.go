package main

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// 初始化数据库，创建对应数据表
func main() {
	common.SetupConfig()
	common.SetupDB()
	tx := common.AuroraRW.Begin()
	var count int

	models := []interface{}{
		&model.Article{},
		&model.Author{},
		&model.Tag{},
		&model.Category{},
		&model.Site{},
	}

	defer func() {
		_ = tx.Close()
	}()

	if err := tx.AutoMigrate(models...).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// 填充默认用户信息
	tx.Model(model.Author{}).Count(&count)
	if count == 0 {
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
	}

	// 填充默认站点信息
	tx.Model(model.Site{}).Count(&count)
	if count == 0 {
		if err := tx.Create(&model.Site{
			Name:    "Skyscraper",
			Powered: "Wimi",
			StartAt: time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()
}
