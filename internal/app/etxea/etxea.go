package etxea

import (
	"github.com/chrisdoherty4/xavier/internal/pkg/etxea"
	"github.com/chrisdoherty4/xavier/internal/pkg/xavier"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea plugin.
func Run(args []string) error {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: xavier.Handshake,
		Plugins:         etxea.Plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
		Logger:          hclog.Default(),
	})
	return nil
}
