package transport

import (
	"context"
	"net/url"
)

// Server ...
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

// Endpointer ...
type Endpointer interface {
	Endpoint() (*url.URL, error)
}
