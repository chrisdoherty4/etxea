package etxea

import (
	"github.com/chrisdoherty4/xavier/xavier/pkg/xavier"
	"github.com/hashicorp/go-plugin"
)

func NewServeConfig() *plugin.ServeConfig {
	return &plugin.ServeConfig{
		HandshakeConfig: xavier.Handshake,
		Plugins:         xavier.Plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
	}
}
