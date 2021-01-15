package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
	"github.com/jinzhu/gorm"
)

type ArticleService struct {
}

func (*ArticleService) GetArticles(params map[string]interface{}, size int, page int) ([]model.Article, error) {
	articleList := make([]model.Article, 0)
	scopes := make([]func(*gorm.DB) *gorm.DB, 0)

	// Todo params 限制条件

	if err := common.AuroraR.Model(model.Article{}).Scopes(scopes...).Limit(size).Offset(size * (page - 1)).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func (*ArticleService) Count(params map[string]interface{}) (int, error) {
	sql := "SELECT COUNT(*) AS count FROM (SELECT * FROM article WHERE 1=1"
	type Result struct {
		Count int
	}
	var queryParams []interface{}
	result := &Result{}

	// Todo params 限制条件

	sql += ") AS t"

	if err := common.AuroraR.Raw(sql, queryParams...).Scan(result).Error; err != nil {
		return -1, err
	}
	return result.Count, nil
}

func (*ArticleService) CreateArticle(article *model.Article) (int, error) {
	if err := common.AuroraRW.Model(model.Article{}).Create(article).Error; err != nil {
		return -1, err
	}
	return article.ID, nil
}

func (*ArticleService) UpdateArticle(articleId int, article map[string]interface{}) error {
	if err := common.AuroraRW.Model(model.Article{}).Where("id = ?", articleId).Update(article).Error; err != nil {
		return err
	}
	return nil
}

func (*ArticleService) DeleteArticle(articleId int) error {
	if err := common.AuroraRW.Model(model.Article{}).Where("id = ?", articleId).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	return nil
}
