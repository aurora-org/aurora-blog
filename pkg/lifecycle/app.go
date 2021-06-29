package lifecycle

import (
	"context"
	"github.com/google/uuid"
	"log"
	"os"
	"syscall"
)

// AppInfo application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	//servers []*http.Server
	instance *ServiceInstance
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		logger: log.Default(),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   options,
	}
}

// ID returns app instance id.
func (a *App) ID() string { return a.opts.id }

// Name returns service name.
func (a *App) Name() string { return a.opts.name }

// Version returns app version.
func (a *App) Version() string { return a.opts.version }

// Metadata returns service metadata.
func (a *App) Metadata() map[string]string { return a.opts.metadata }

// Endpoint returns endpoints.
func (a *App) Endpoint() []string { return a.instance.Endpoint }

// Run ...
func (a *App) Run() error {
	// Todo
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	// Todo
	return nil
}