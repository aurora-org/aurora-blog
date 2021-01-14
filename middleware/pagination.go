package middleware

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/tool"
	"github.com/kataras/iris/v12"
)

type Pagination struct{}

// 分页参数校验，size 最大为500
func (*Pagination) Limit(ctx iris.Context) {
	page := tool.StringToInt(ctx.URLParam("page"))
	size := tool.StringToInt(ctx.URLParam("size"))
	if page < 1 || size < 1 || size > 500 {
		common.Render(ctx, status.BadRequest, "invalid pagination")
		return
	}
	ctx.Next()
}
