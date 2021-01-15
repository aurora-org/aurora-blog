package mng

import (
	ctrls "aurora/blog/api/controller"
	"github.com/kataras/iris/v12/core/router"
)

func SetupSite(v router.Party) {
	var (
		siteController = ctrls.SiteController{}
	)

	v.Get("/site", siteController.GetSiteInfo)
	v.Post("/site", siteController.UpdateSiteInfo)
}
