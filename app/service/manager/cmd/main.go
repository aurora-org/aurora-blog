package main

import "aurora/blog/api/pkg/lifecycle"

func main() {
	app := lifecycle.New()
	_ = app.Run()
}
