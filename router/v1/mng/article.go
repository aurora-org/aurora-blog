package mng

import (
	ctrls "aurora/blog/api/controller"
	"github.com/kataras/iris/v12/core/router"
)

func SetupArticle(v router.Party) {

	var (
		articleController = ctrls.ArticleController{}
		//page           = middleware.Pagination{}
	)

	v.Get("/articles", articleController.GetArticles)
}
