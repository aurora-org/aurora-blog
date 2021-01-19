package mng

import "github.com/kataras/iris/v12/core/router"

func Setup(v router.Party) {
	api := v.Party("/mng")

	SetupArticle(api)
	SetupAuthor(api)
	SetupCategory(api)
	SetupTag(api)
	SetupSite(api)
	SetupAccount(api)
}
