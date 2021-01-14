package common

import (
	"aurora/blog/api/constant/status"
	"github.com/kataras/iris/v12"
)

type Result struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (*Result) NewResult(code int, status string, data interface{}) *Result {
	return &Result{
		Code:   code,
		Status: status,
		Data:   data,
	}
}

func Render(ctx iris.Context, code int, data interface{}) {
	result := &Result{}
	_, _ = ctx.JSON(result.NewResult(code, status.CodeMap[code], data))
}
