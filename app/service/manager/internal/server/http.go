package server

import (
	"aurora/blog/api/pkg/lifecycle"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpServer struct {
	engine *gin.Engine
	*http.Server
}

func (h *HttpServer) Start(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)
	fmt.Println(info.Name(), "start")
	return nil
}

func (h *HttpServer) Stop(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)
	fmt.Println(info.Name(), "stop")
	return nil
}
