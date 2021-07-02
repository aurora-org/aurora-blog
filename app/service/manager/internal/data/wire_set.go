package data

import (
	"aurora/blog/api/pkg/config"
	"context"
	"entgo.io/ent/examples/start/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProvideSet for data package ...
var ProvideSet = wire.NewSet(NewData, NewUserRepo)

// NewData return new *gorm.DB wrapper.
func NewData(logger *zap.Logger, config config.Data) (*Data, func(), error) {
	// open database
	client, err := ent.Open(config.Driver(), config.Source())
	if err != nil {
		return nil, nil, err
	}
	// create data schema
	logger.Sugar().Info("[AURORA] create table")
	err = client.Schema.Create(context.Background())
	if err != nil {
		return nil, nil, err
	}

	return &Data{db: client}, func() {
		logger.Sugar().Info("[AURORA] cleanup db")
		_ = client.Close()
	}, nil
}

type Data struct {
	db *ent.Client
}
