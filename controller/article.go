package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant/article"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/entity"
	"aurora/blog/api/model"
	"aurora/blog/api/service"
	"aurora/blog/api/tool"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
	"strings"
	"sync"
)

type ArticleController struct {
	mutex sync.Mutex
}

func (self *ArticleController) ReadArticle(ctx iris.Context) {
	articleService := service.ArticleService{}
	articleId, _ := ctx.Params().GetInt("articleId")

	self.mutex.Lock()
	articleService.ReadArticle(articleId)
	self.mutex.Unlock()
}

func (*ArticleController) GetArticles(ctx iris.Context) {
	articleService := service.ArticleService{}

	size := tool.StringToInt(ctx.URLParam("size"))
	page := tool.StringToInt(ctx.URLParam("page"))

	tag := ctx.URLParam("tag")
	category := ctx.URLParam("category")
	order := ctx.URLParam("order")
	by := ctx.URLParam("by")

	queryParams := map[string]interface{}{}

	if tag != "" {
		queryParams["tag"] = tag
	}
	if category != "" {
		queryParams["category"] = category
	}

	if order != "" {
		queryParams["order"] = article.OrderMap[order]
	} else {
		queryParams["order"] = article.OrderMap["CREATE"]
	}
	if by != "" {
		queryParams["by"] = strings.ToUpper(by)
	} else {
		queryParams["by"] = "DESC"
	}

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

func (*ArticleController) GetArticle(ctx iris.Context) {
	articleService := service.ArticleService{}
	articleId, _ := ctx.Params().GetInt("articleId")

	article, err := articleService.GetArticleById(articleId)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.OK, article.Mapping())
}

func (*ArticleController) CreateArticle(ctx iris.Context) {
	authorService := service.AuthorService{}
	articleService := service.ArticleService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	author, err := authorService.GetFirstAuthor()
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	article := &model.Article{
		Title:    params.Get("title").MustString(),
		Content:  params.Get("content").MustString(),
		Banner:   params.Get("banner").MustString(),
		Author:   author.NickName,
		Tag:      params.Get("tag").MustString(),
		Category: params.Get("category").MustString(),
		Times:    0,
		Visible:  params.Get("visible").MustBool(),
		Extra:    params.Get("extra").MustString(),
	}

	articleId, err := articleService.CreateArticle(article)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	saved, err := articleService.GetArticleById(articleId)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.Created, saved.Mapping())
}

func (*ArticleController) UpdateArticle(ctx iris.Context) {
	articleService := service.ArticleService{}
	articleId, _ := ctx.Params().GetInt("articleId")

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

	if err = articleService.UpdateArticle(articleId, paramsMap); err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	updated, err := articleService.GetArticleById(articleId)
	if err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.Updated, updated.Mapping())
}

func (*ArticleController) DeleteArticle(ctx iris.Context) {
	articleService := service.ArticleService{}
	articleId, _ := ctx.Params().GetInt("articleId")

	if err := articleService.DeleteArticle(articleId); err != nil {
		common.Render(ctx, status.InternalServerError, err.Error())
		return
	}

	common.Render(ctx, status.Deleted, articleId)
}
