package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/model"
	"aurora/blog/api/service"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

type CategoryController struct {
}

func (*CategoryController) GetCategories(ctx iris.Context) {
	categoryService := service.CategoryService{}

	categories, err := categoryService.GetCategories()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	categoryList := make([]map[string]interface{}, 0)
	for _, item := range categories {
		categoryList = append(categoryList, item.Mapping())
	}

	common.Render(ctx, status.OK, categoryList)
}

func (*CategoryController) CreateCategory(ctx iris.Context) {
	categoryService := service.CategoryService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	category := &model.Category{
		Name:        params.Get("name").MustString(),
		Description: params.Get("description").MustString(),
		Extra:       params.Get("extra").MustString(),
	}

	categoryId, err := categoryService.CreateCategory(category)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	saved, err := categoryService.GetCategoryById(categoryId)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.Created, saved.Mapping())
}

func (*CategoryController) DeleteCategory(ctx iris.Context) {
	categoryService := service.CategoryService{}
	categoryId, _ := ctx.Params().GetInt("categoryId")

	if err := categoryService.DeleteCategory(categoryId); err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.Deleted, categoryId)
}
