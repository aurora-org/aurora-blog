package mng

import (
	ctrls "aurora/blog/api/controller"
	"github.com/kataras/iris/v12/core/router"
)

func SetupTag(v router.Party) {
	var (
		tagController = ctrls.TagController{}
	)

	v.Get("/tags", tagController.GetTags)
	v.Post("/tags", tagController.CreateTag)
	v.Delete("/tags", tagController.DeleteTag)
}
