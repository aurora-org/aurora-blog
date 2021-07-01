package http

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/log"
	"aurora/blog/api/pkg/transport"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server http server with gin.Engine
type Server struct {
	*http.Server

	Metadata map[string]string
	err      error
	log      *log.Logger
}

func NewServer(engine *gin.Engine) transport.Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":8080", // Todo to be set by config
			Handler: engine,
		},
		Metadata: make(map[string]string),
		log:      log.NewLogger(),
		err:      nil,
	}
}

func (s *Server) Start(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)
	s.log.Println(info.Name(), "start")
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)
	s.log.Println(info.Name(), "stop")
	return s.Shutdown(ctx)
}
