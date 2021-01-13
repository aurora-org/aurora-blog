package common

import (
	"github.com/kataras/iris/v12"
	"runtime/debug"
)

func NewPanicHandler(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
