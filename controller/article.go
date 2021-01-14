package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"github.com/kataras/iris/v12"
)

type ArticleController struct {
}

func (*ArticleController) GetArticles(ctx iris.Context) {
	common.Render(ctx, status.OK, "hello world")
}
