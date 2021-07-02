package data

import (
	"github.com/google/wire"
)

// ProvideSet for data package ...
var ProvideSet = wire.NewSet(NewData, NewHelloRepo)

