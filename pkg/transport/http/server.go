package http

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/transport"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

var _ transport.Server = (*Server)(nil)

// Server http server with *gin.Engine
type Server struct {
	*http.Server

	Metadata map[string]string
	log      *zap.Logger
}

func (s *Server) Endpoint() (*url.URL, error) {
	return &url.URL{
		Host:        "127.0.0.1:8080",
		Path:        "/",
	}, nil
}

func NewServer(engine *gin.Engine, logger *zap.Logger) transport.Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":8080", // Todo read from config
			Handler: engine,
		},
		Metadata: make(map[string]string),
		log:      logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	// Todo endpoint error
	//fmt.Println(info.Endpoint())
	s.log.Info(fmt.Sprintf("[AURORA] server start"),
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.Any("metadata", info.Metadata()))
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)
	s.log.Info(fmt.Sprintf("[AURORA] server stop"),
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.Any("metadata", info.Metadata()))
	return s.Shutdown(ctx)
}
