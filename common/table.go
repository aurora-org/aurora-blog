package common

import (
	"aurora/blog/api/model"
	"aurora/blog/api/tool"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func SetupTable() {
	tx := AuroraRW.Begin()
	var count int

	models := []interface{}{
		&model.Article{},
		&model.Author{},
		&model.Tag{},
		&model.Category{},
		&model.Site{},
		&model.Account{},
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

	// 填充默认账户名和密码
	tx.Model(model.Account{}).Count(&count)
	if count == 0 {
		if err := tx.Create(&model.Account{
			UserName: "aurora",
			Password: tool.Md5Encode("aurora"),
		}).Error; err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()
}
