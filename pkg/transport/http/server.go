package http

import (
	"aurora/blog/api/pkg/config"
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/transport"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

var _ transport.Server = (*Server)(nil)

// Server http server with *gin.Engine
type Server struct {
	*http.Server

	addr *url.URL
	log  *zap.Logger
}

func (s *Server) Endpoint() *url.URL {
	return s.addr
}

func NewServer(engine *gin.Engine, logger *zap.Logger, config config.Server) transport.Server {
	return &Server{
		Server: &http.Server{
			Addr:         config.Addr(),
			Handler:      engine,
			WriteTimeout: config.Timeout(),
		},
		addr: &url.URL{
			Host: config.Addr(),
		},
		log: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("[AURORA] server start",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.Strings("endpoints", info.Endpoint()),
		zap.Any("metadata", info.Metadata()))

	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("[AURORA] server stop",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.Strings("endpoints", info.Endpoint()),
		zap.Any("metadata", info.Metadata()))

	return s.Shutdown(ctx)
}
