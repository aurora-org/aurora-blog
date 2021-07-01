package conf

import (
	"aurora/blog/api/pkg/log"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProvideSet for data package ...
var ProvideSet = wire.NewSet(NewLogger)

func NewLogger() *zap.Logger {
	return log.NewLogger("aurora-manager")
}
