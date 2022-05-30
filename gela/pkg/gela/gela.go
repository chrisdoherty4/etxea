package gela

import (
	"github.com/chrisdoherty4/etxea/pkg/etxea"
	"github.com/hashicorp/go-plugin"
)

func NewServeConfig() *plugin.ServeConfig {
	return &plugin.ServeConfig{
		HandshakeConfig: etxea.Handshake,
		Plugins:         etxea.Plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
	}
}
