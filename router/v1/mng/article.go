package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupArticle(v router.Party) {

	var (
		articleController = ctrls.ArticleController{}
		authorization     = middleware.Authorization{}
		pagination        = middleware.Pagination{}
	)

	v.Get("/articles", pagination.Limit, articleController.GetArticles)
	v.Post("/articles", authorization.CheckToken, articleController.CreateArticle)
	v.Put("/articles/{articleId:int}", authorization.CheckToken, articleController.UpdateArticle)
	v.Delete("/articles/{articleId:int}", authorization.CheckToken, articleController.DeleteArticle)
}
