package middleware

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/tool"
	"github.com/kataras/iris/v12"
)

type Authorization struct {
}

func (*Authorization) CheckToken(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		common.Render(ctx, status.Forbidden, "token missing")
		return
	}

	claim, err := tool.ParseToken(token)
	if err != nil {
		common.Render(ctx, status.Forbidden, err.Error())
		return
	}

	ctx.Values().Set("claim", claim)
	ctx.Next()
}
