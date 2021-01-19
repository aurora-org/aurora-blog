package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupTag(v router.Party) {
	var (
		tagController = ctrls.TagController{}
		authorization = middleware.Authorization{}
	)

	v.Get("/tags", tagController.GetTags)
	v.Post("/tags", tagController.CreateTag)
	v.Delete("/tags/{tagId:int}", authorization.CheckToken, tagController.DeleteTag)
}
