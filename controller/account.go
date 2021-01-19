package controller

import (
	"aurora/blog/api/common"
	"aurora/blog/api/constant"
	"aurora/blog/api/constant/status"
	"aurora/blog/api/entity"
	"aurora/blog/api/model"
	"aurora/blog/api/service"
	"aurora/blog/api/tool"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

type AccountController struct {
}

func (*AccountController) Login(ctx iris.Context) {
	accountService := service.AccountService{}

	var token string
	var err error

	params := simplejson.New()
	if err = ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	account := &model.Account{
		UserName: params.Get("userName").MustString(),
		Password: params.Get("password").MustString(),
	}

	if accountService.CheckAccount(account) {
		token, err = tool.GenerateToken(account.UserName)
		if err != nil {
			common.Render(ctx, status.InternalServerError, err.Error())
			return
		}
	} else {
		common.Render(ctx, status.BadRequest, "password error")
		return
	}

	common.Render(ctx, status.OK, map[string]interface{}{
		"token": token,
	})
}

func (*AccountController) ChangePassword(ctx iris.Context) {
	accountService := service.AccountService{}

	params := simplejson.New()
	if err := ctx.ReadJSON(&params); err != nil {
		common.Render(ctx, status.BadRequest, err.Error())
		return
	}

	oldPwd := params.Get("old").MustString()
	newPwd := params.Get("new").MustString()
	account := &model.Account{
		Password: oldPwd,
	}

	// 获取中间件中解析的 claim
	if claim, ok := ctx.Values().Get(constant.ContextClaimKey).(*entity.CustomClaim); ok {
		account.UserName = claim.UserName
	}

	if accountService.CheckAccount(account) {
		if err := accountService.UpdatePassword(newPwd); err != nil {
			common.Render(ctx, status.InternalServerError, err.Error())
			return
		}
	} else {
		common.Render(ctx, status.BadRequest, "password error")
		return
	}

	common.Render(ctx, status.Updated, nil)
}
