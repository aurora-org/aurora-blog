package data

import (
	"aurora/blog/api/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Data struct {
	db *gorm.DB
}

// NewData return new *gorm.DB wrapper.
func NewData(logger *zap.Logger, data config.Data) (*Data, func(), error) {
	db, err := gorm.Open(data.Driver(), data.Source())
	if err != nil {
		return nil, nil, err
	}
	return &Data{db: db}, func() {
		logger.Sugar().Info("[AURORA] cleanup db")
		_ = db.Close()
	}, nil
}
