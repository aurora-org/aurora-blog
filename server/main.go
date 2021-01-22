package server

import (
	"aurora/blog/api/common"
	"aurora/blog/api/router"
	"github.com/kataras/iris/v12"
	"log"
)

func SetupServer() {
	app := iris.New()
	app.Use(common.NewPanicHandler)
	router.Route(app)
	_ = app.Run(iris.Addr(":8080"),
		iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
}

func SetupGraphQL() {
	app := iris.New()
	app.Use(common.NewPanicHandler)
	// Todo graphql api
	//app.Get("/graphql")

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("---- graphql server error ----")
				log.Println("Recovery_error: ", err)
			}
		}()
		_ = app.Run(iris.Addr(":8888"),
			iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
	}()
}
