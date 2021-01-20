package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupAuthor(v router.Party) {
	var (
		authorController = ctrls.AuthorController{}
		authorization    = middleware.Authorization{}
	)

	v.Get("/author", authorController.GetAuthorInfo)
	v.Put("/author", authorization.CheckToken, authorController.UpdateAuthorInfo)
}
