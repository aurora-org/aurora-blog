package service

import (
	"aurora/blog/api/common"
	"aurora/blog/api/model"
	"aurora/blog/api/tool"
)

type AccountService struct {
}

func (*AccountService) CheckAccount(account *model.Account) bool {
	var count int
	common.AuroraR.Model(model.Account{}).Where("user_name = ?", account.UserName).Where("password = ?", tool.Md5Encode(account.Password)).Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}

func (*AccountService) UpdatePassword(password string) error {
	if err := common.AuroraRW.Model(model.Account{}).Update("password", tool.Md5Encode(password)).Error; err != nil {
		return err
	}
	return nil
}
