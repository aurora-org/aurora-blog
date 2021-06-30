package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _ biz.Repo = (*Data)(nil)

type Data struct {
	db *gorm.DB
}

func (d *Data) Hello() (string, error) {
	// call db
	return "hello world", nil
}
