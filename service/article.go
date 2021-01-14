package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
)

type ArticleService struct {
}

func (*ArticleService) GetAll(params map[string]interface{}, size int, page int) ([]model.Article, error) {
	var articleList []model.Article
	if err := common.AuroraR.Limit(size).Offset(size * (page - 1)).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func (*ArticleService) Create(article *model.Article) (*model.Article, error) {
	if err := common.AuroraRW.Model(model.Article{}).Create(article).Error; err != nil {

	}
	return nil, nil
}

func (*ArticleService) Update(article map[string]interface{}) (*model.Article, error) {
	return nil, nil
}
