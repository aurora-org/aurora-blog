package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
)

type SiteService struct {
}

func (*SiteService) GetFirstSite() (*model.Site, error) {
	site := &model.Site{}
	if err := common.AuroraR.Model(model.Site{}).First(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}

func (*SiteService) UpdateFirstSite(site map[string]interface{}) error {
	if err := common.AuroraRW.Model(model.Site{}).Update(site).Error; err != nil {
		return err
	}
	return nil
}
