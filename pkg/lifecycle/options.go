package lifecycle

import (
	"context"
	"log"
	"net/url"
	"os"
)

// Option application option.
type Option func(*options)

// options application option struct.
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx  context.Context
	sigs []os.Signal

	logger    *log.Logger
	registrar Registrar
	servers   []Server
}

// WithID with service id.
func WithID(id string) Option {
	return func(o *options) { o.id = id }
}

// WithName with service name.
func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

// WithVersion with service version.
func WithVersion(version string) Option {
	return func(o *options) { o.version = version }
}

// WithMetadata with service metadata.
func WithMetadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// WithEndpoint with service endpoint.
func WithEndpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// WithContext with service context.
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithLogger with service logger.
func WithLogger(logger *log.Logger) Option {
	return func(o *options) { o.logger = logger }
}

// WithSignal with exit signals.
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// WithRegistrar with service registry.
func WithRegistrar(r Registrar) Option {
	return func(o *options) { o.registrar = r }
}

// WithServer with transport servers.
func WithServer(srv ...Server) Option {
	return func(o *options) { o.servers = srv }
}
