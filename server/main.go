package server

import (
	"aurora/blog/api/common"
	"aurora/blog/api/router"
	"github.com/kataras/iris/v12"
)

func InitServer() {
	app := iris.New()
	app.Use(common.NewPanicHandler)
	router.Route(app)
	_ = app.Run(iris.Addr(":8080"),
		iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
}
