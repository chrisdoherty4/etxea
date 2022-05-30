package gela

import (
	"github.com/chrisdoherty4/etxea/gela/pkg/gela"
	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea plugin.
func Run(args []string) error {
	plugin.Serve(gela.NewServeConfig())
	return nil
}
