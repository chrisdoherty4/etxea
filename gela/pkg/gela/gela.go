package gela

import (
	"github.com/chrisdoherty4/etxea/pkg/etxea"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewServeConfig(plugins plugin.PluginSet) *plugin.ServeConfig {
	return &plugin.ServeConfig{
		HandshakeConfig: etxea.Handshake,
		Plugins:         plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
		Logger:          hclog.Default(),
	}
}
