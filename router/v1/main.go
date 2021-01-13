package v1

import (
	mngApi "aurora/blog/api/router/v1/mng"
	"github.com/kataras/iris/v12"
)

func SetupApi(app *iris.Application) {
	api := app.Party("/api/v1")

	mngApi.Setup(api)
}
