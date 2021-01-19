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
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.OK, site.Mapping())
}

func (*SiteController) UpdateSiteInfo(ctx iris.Context) {
	siteService := service.SiteService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	paramsMap, err := params.Map()
	if err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	if err = siteService.UpdateFirstSite(paramsMap); err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	updated, err := siteService.GetFirstSite()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}
	common.Render(ctx, status.Updated, updated.Mapping())
}
