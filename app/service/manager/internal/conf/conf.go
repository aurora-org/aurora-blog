package conf

import (
	"aurora/blog/api/pkg/config"
	"time"
)

var _ config.Server = (*ServerConfig)(nil)
var _ config.Data = (*DataConfig)(nil)

// ServerConfig ...
type ServerConfig struct {
	addr    string
	timeout time.Duration
}

func (c *ServerConfig) Addr() string {
	return c.addr
}

func (c *ServerConfig) Timeout() time.Duration {
	return c.timeout
}

func NewServerConfig() config.Server {
	return &ServerConfig{
		addr:    "127.0.0.1:8080",
		timeout: time.Second,
	}
}

// DataConfig ...
type DataConfig struct {
	driver string
	source string
}

func (d *DataConfig) Driver() string {
	return d.driver
}

func (d *DataConfig) Source() string {
	return d.source
}

func NewDataConfig() config.Data {
	return &DataConfig{
		driver: "mysql",
		source: "root:123@tcp(127.0.0.1:3306)/hello?parseTime=True",
	}
}
