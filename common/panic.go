package common

import (
	"aurora/blog/api/constant/status"
	"github.com/kataras/iris/v12"
	"runtime/debug"
)

func NewPanicHandler(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			ctx.StatusCode(status.InternalServerError)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
