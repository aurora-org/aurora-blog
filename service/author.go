package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
)

type AuthorService struct {
}

func (*AuthorService) GetFirstAuthor() (*model.Author, error) {
	author := &model.Author{}
	if err := common.AuroraR.Model(model.Author{}).First(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (*AuthorService) UpdateFirstAuthor(author map[string]interface{}) error {
	if err := common.AuroraRW.Model(model.Author{}).Update(author).Error; err != nil {
		return err
	}
	return nil
}
