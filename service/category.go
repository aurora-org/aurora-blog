package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
)

type CategoryService struct {
}

func (*CategoryService) GetCategories() ([]model.Category, error) {
	categoryList := make([]model.Category, 0)
	if err := common.AuroraR.Model(model.Category{}).Find(&categoryList).Error; err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (*CategoryService) GetCategoryById(categoryId int) (*model.Category, error) {
	category := &model.Category{}
	if err := common.AuroraR.Model(model.Category{}).Where("id = ?", categoryId).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (*CategoryService) CreateCategory(category *model.Category) (int, error) {
	if err := common.AuroraRW.Model(model.Category{}).Create(category).Error; err != nil {
		return -1, err
	}
	return category.ID, nil
}

func (*CategoryService) DeleteCategory(categoryId int) error {
	if err := common.AuroraRW.Model(model.Category{}).Where("id = ?", categoryId).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}

func (*CategoryService) IsExistCategory(categoryName string) bool {
	var count int
	common.AuroraR.Model(model.Category{}).Where("name = ?", categoryName).Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}
