package lifecycle

import (
	"aurora/blog/api/pkg/registry"
	"aurora/blog/api/pkg/transport"
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// AppInfo application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

// App application struct.
type App struct {
	opts     options
	ctx      context.Context
	cancel   func()
	instance *registry.ServiceInstance
}

// New return new App
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

// Run start application.
func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}
	ctx := NewContext(a.ctx, a)
	group, ctx := errgroup.WithContext(ctx)
	wg := &sync.WaitGroup{}

	// start servers
	for _, srv := range a.opts.servers {
		group.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		wg.Add(1)
		group.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	// waiting for app server start
	wg.Wait()

	// register instance.
	if a.opts.registrar != nil {
		if err := a.opts.registrar.Register(a.opts.ctx, instance); err != nil {
			return err
		}
		a.instance = instance
	}

	// wait for os signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, a.opts.sigs...)
	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sigs:
				return a.Stop()
			}
		}
	})
	if err := group.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	// deregister instance
	if a.opts.registrar != nil && a.instance != nil {
		if err := a.opts.registrar.Deregister(a.opts.ctx, a.instance); err != nil {
			return err
		}
	}
	// cancel context
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

// buildInstance build ServiceInstance with options.endpoints
func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	var endpoints []string
	for _, e := range a.opts.endpoints {
		endpoints = append(endpoints, e.String())
	}
	if len(endpoints) == 0 {
		for _, srv := range a.opts.servers {
			if r, ok := srv.(transport.Endpointer); ok {
				e, err := r.Endpoint()
				if err != nil {
					return nil, err
				}
				endpoints = append(endpoints, e.String())
			}
		}
	}
	return &registry.ServiceInstance{
		ID:       a.opts.id,
		Name:     a.opts.name,
		Version:  a.opts.version,
		Metadata: a.opts.metadata,
		Endpoint: endpoints,
	}, nil
}

// appKey app context key
type appKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}
