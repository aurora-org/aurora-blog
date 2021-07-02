package config

import "time"

type Server interface {
	Addr() string
	Timeout() time.Duration
}

type Data interface {
	Driver() string
	Source() string
}
