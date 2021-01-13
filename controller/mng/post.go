package mng

import "github.com/kataras/iris/v12"

type PostController struct {
}

func (*PostController) GetPosts(ctx iris.Context) {
	_, _ = ctx.Write([]byte("hello world"))
}
