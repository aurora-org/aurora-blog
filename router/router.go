package router

import (
	v "aurora/blog/api/router/v1"
	"github.com/kataras/iris/v12"
)

func Route(app *iris.Application) {
	v.SetupApi(app)
}
