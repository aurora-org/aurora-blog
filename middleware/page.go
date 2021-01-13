package middleware

import (
	"aurora/blog/api/tool"
	"github.com/kataras/iris/v12"
)

type Pagination struct{}

// 分页参数校验
func (*Pagination) Limit(ctx iris.Context) {
	limit := tool.StringToInt(ctx.URLParam("limit"))
	if limit > 100 || limit <= 0 {
		// Todo
		//_, _ = ctx.JSON()
		return
	}
	ctx.Next()
}
