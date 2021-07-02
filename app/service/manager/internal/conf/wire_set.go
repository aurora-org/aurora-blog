package conf

import (
	"aurora/blog/api/pkg/log"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProvideSet for conf package ...
var ProvideSet = wire.NewSet(NewZapLogger, NewServerConfig, NewDataConfig)

// NewZapLogger return new *zap.Logger
func NewZapLogger() *zap.Logger {
	return log.NewLogger("aurora-manager")
}
