package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/service"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

type SiteController struct {
}

func (*SiteController) GetSiteInfo(ctx iris.Context) {
	siteService := service.SiteService{}

	site, err := siteService.GetFirstSite()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	common.Render(ctx, status.OK, site.Mapping())
}

func (*SiteController) UpdateSiteInfo(ctx iris.Context) {
	siteService := service.SiteService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err)
		return
	}

	paramsMap, err := params.Map()
	if err != nil {
		common.Render(ctx, status.BadRequest, err)
		return
	}

	if err = siteService.UpdateFirstSite(paramsMap); err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	updated, err := siteService.GetFirstSite()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}
	common.Render(ctx, status.Updated, updated.Mapping())
}
