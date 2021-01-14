package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/mustache"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func _setupDB(name string, readOnly bool) *gorm.DB {
	isDevelopment := viper.GetBool("isDevelopment")
	readWrite := "wr"
	if readOnly {
		readWrite = "r"
	}

	prefix := mustache.Render("config.database.{{name}}.{{readWrite}}", map[string]interface{}{
		"name":      name,
		"readWrite": readWrite,
	})

	dialect := viper.GetString(prefix + ".dialect")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetString(prefix + ".port")

	url := mustache.Render("{{user}}:{{password}}@tcp({{host}}:{{port}})/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
		"user":     user,
		"password": password,
		"database": database,
		"host":     host,
		"port":     port,
	})

	db, err := gorm.Open(dialect, url)
	if err != nil {
		panic(err)
	}

	db.LogMode(isDevelopment)
	return db
}

func SetupDB() {
	AuroraRW = _setupDB("aurora-db", false)
	AuroraR = _setupDB("aurora-db", true)
}

var (
	AuroraRW *gorm.DB
	AuroraR  *gorm.DB
)
