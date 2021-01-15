package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/service"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

type AuthorController struct {
}

func (*AuthorController) GetAuthorInfo(ctx iris.Context) {
	authorService := service.AuthorService{}

	author, err := authorService.GetFirstAuthor()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	common.Render(ctx, status.OK, author.Mapping())
}

func (*AuthorController) UpdateAuthorInfo(ctx iris.Context) {
	authorService := service.AuthorService{}

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

	if err = authorService.UpdateFirstAuthor(paramsMap); err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}

	updated, err := authorService.GetFirstAuthor()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err)
		return
	}
	common.Render(ctx, status.Updated, updated.Mapping())
}
