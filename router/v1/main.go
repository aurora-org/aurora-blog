package v1

import (
	mngApi "aurora/blog/api/router/v1/mng"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func SetupApi(app *iris.Application) {
	api := app.Party("/api/v1", cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})).AllowMethods(iris.MethodOptions)

	mngApi.Setup(api)
}
