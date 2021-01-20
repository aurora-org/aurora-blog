package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
)

type TagService struct {
}

func (*TagService) GetTags() ([]model.Tag, error) {
	tagList := make([]model.Tag, 0)
	if err := common.AuroraR.Model(model.Tag{}).Find(&tagList).Error; err != nil {
		return nil, err
	}
	return tagList, nil
}

func (*TagService) GetTagById(tagId int) (*model.Tag, error) {
	tag := &model.Tag{}
	if err := common.AuroraR.Model(model.Tag{}).Where("id = ?", tagId).First(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (*TagService) CreateTag(tag *model.Tag) (int, error) {
	if err := common.AuroraRW.Model(model.Tag{}).Create(tag).Error; err != nil {
		return -1, err
	}
	return tag.ID, nil
}

func (*TagService) DeleteTag(tagId int) error {
	if err := common.AuroraRW.Model(model.Tag{}).Where("id = ?", tagId).Delete(&model.Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func (*TagService) IsExistTag(tagName string) bool {
	var count int
	common.AuroraR.Model(model.Tag{}).Where("name = ?", tagName).Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}
