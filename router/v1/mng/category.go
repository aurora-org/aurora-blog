package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupCategory(v router.Party) {
	var (
		categoryController = ctrls.CategoryController{}
		authorization      = middleware.Authorization{}
	)

	v.Get("/categories", categoryController.GetCategories)
	v.Post("/categories", authorization.CheckToken, categoryController.CreateCategory)
	v.Delete("/categories/{categoryId:int}", authorization.CheckToken, categoryController.DeleteCategory)
}
