package mng

import (
	ctrls "aurora/blog/api/controller"
	"github.com/kataras/iris/v12/core/router"
)

func SetupAuthor(v router.Party) {
	var (
		authorController = ctrls.AuthorController{}
	)

	v.Get("/author", authorController.GetAuthorInfo)
}
