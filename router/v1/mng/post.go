package mng

import (
	mngCtrls "aurora/blog/api/controller/mng"
	"github.com/kataras/iris/v12/core/router"
)

func SetupPost(v router.Party) {

	var (
		postController = mngCtrls.PostController{}
		//page           = middleware.Pagination{}
	)

	v.Get("/posts", postController.GetPosts)
}
