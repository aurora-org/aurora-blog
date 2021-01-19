package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/entity"
	"aurora/blog/api/model"
	"aurora/blog/api/service"
	"aurora/blog/api/tool"
	"github.com/kataras/iris/v12"
)

type ArticleController struct {
}

func (*ArticleController) GetArticles(ctx iris.Context) {
	articleService := service.ArticleService{}
	// Todo 接收过滤参数
	size := tool.StringToInt(ctx.URLParam("size"))
	page := tool.StringToInt(ctx.URLParam("page"))

	queryParams := map[string]interface{}{}

	count, err := articleService.Count(queryParams)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	articleList := make([]map[string]interface{}, 0)
	articles := make([]model.Article, 0)
	if count > 0 {
		articles, err = articleService.GetArticles(queryParams, size, page)
		if err != nil {
			common.Render(ctx, status.InternalServerError, err.Error())
			return
		}

		for _, item := range articles {
			articleList = append(articleList, item.Mapping())
		}
	}

	common.Render(ctx, status.OK, entity.PaginationData{
		Pagination: entity.Pagination{
			Total: count,
			Page:  page,
			Size:  size,
		},
		Objects: articleList,
	})
}

func (*ArticleController) CreateArticle(ctx iris.Context) {

}

func (*ArticleController) UpdateArticle(ctx iris.Context) {

}

func (*ArticleController) DeleteArticle(ctx iris.Context) {

}
