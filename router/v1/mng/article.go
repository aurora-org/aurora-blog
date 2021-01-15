package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupArticle(v router.Party) {

	var (
		articleController = ctrls.ArticleController{}
		pagination        = middleware.Pagination{}
	)

	v.Get("/articles", pagination.Limit, articleController.GetArticles)
}
