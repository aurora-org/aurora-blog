package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/model"
	"aurora/blog/api/service"
	"aurora/blog/api/tool"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

type TagController struct {
}

func (*TagController) GetTags(ctx iris.Context) {
	tagService := service.TagService{}

	tags, err := tagService.GetTags()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	tagList := make([]map[string]interface{}, 0)
	for _, item := range tags {
		tagList = append(tagList, item.Mapping())
	}

	common.Render(ctx, status.OK, tagList)
}

func (*TagController) CreateTag(ctx iris.Context) {
	tagService := service.TagService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err)
		return
	}

	tag := &model.Tag{
		Name:  params.Get("name").MustString(),
		Extra: params.Get("extra").MustString(),
	}

	tagId, err := tagService.CreateTag(tag)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	saved, err := tagService.GetTagById(tagId)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
	}

	common.Render(ctx, status.Created, saved.Mapping())
}

func (*TagController) DeleteTag(ctx iris.Context) {
	tagService := service.TagService{}
	tagId := tool.StringToInt(ctx.URLParam("id"))

	if err := tagService.DeleteTag(tagId); err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	common.Render(ctx, status.Deleted, tagId)
}
