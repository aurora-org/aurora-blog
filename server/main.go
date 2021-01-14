package server

import (
	"aurora/blog/api/common"
	"aurora/blog/api/router"
	"github.com/graphql-go/handler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
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
	app.Get("/graphql", func(ctx context.Context) {
		handler.New(&handler.Config{
			// Todo
		})
	})

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
