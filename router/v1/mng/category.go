package mng

import (
	ctrls "aurora/blog/api/controller"
	"github.com/kataras/iris/v12/core/router"
)

func SetupCategory(v router.Party) {
	var (
		categoryController = ctrls.CategoryController{}
	)

	v.Get("/categories", categoryController.GetCategories)
	v.Post("/categories", categoryController.CreateCategory)
	v.Delete("/categories/{categoryId:int}", categoryController.DeleteCategory)
}
