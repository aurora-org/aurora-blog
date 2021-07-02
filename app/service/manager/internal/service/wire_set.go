package service

import (
	"github.com/google/wire"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewHelloService)
