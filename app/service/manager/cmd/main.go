package main

import (
	"aurora/blog/api/pkg/lifecycle"
	"context"
	"fmt"
)

func main() {
	app := lifecycle.New(lifecycle.WithContext(context.Background()))
	fmt.Println(app.ID())
	_ = app.Run()
}
