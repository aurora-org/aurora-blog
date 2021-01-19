package mng

import (
	ctrls "aurora/blog/api/controller"
	"aurora/blog/api/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func SetupAccount(v router.Party) {

	var (
		accountController = ctrls.AccountController{}
		authorization     = middleware.Authorization{}
	)

	v.Post("/account/login", accountController.Login)
	v.Put("/account/password", authorization.CheckToken, accountController.ChangePassword)
}
