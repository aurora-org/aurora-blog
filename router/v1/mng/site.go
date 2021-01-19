package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupSite(v router.Party) {
	var (
		siteController = ctrls.SiteController{}
		authorization  = middleware.Authorization{}
	)

	v.Get("/site", siteController.GetSiteInfo)
	v.Post("/site", authorization.CheckToken, siteController.UpdateSiteInfo)
}
