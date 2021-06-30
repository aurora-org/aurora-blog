package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// ProvideSet for data package ...
var ProvideSet = wire.NewSet(NewData)

func NewData() (biz.Repo, func(), error) {
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/hello?parseTime=True")
	if err != nil {
		return nil, nil, err
	}
	return &Data{db: db}, func() {
		fmt.Println("cleanup db")
		_ = db.Close()
	}, nil
}
