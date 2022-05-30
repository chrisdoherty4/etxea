package etxea

import (
	"github.com/chrisdoherty4/xavier/etxea/pkg/etxea"
	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea plugin.
func Run(args []string) error {
	plugin.Serve(etxea.NewServeConfig())
	return nil
}
