package gela

import "github.com/hashicorp/go-plugin"

// Run executes the etxea plugin.
func Run(args []string) error {
	plugin.Serve(NewServeConfig())
	return nil
}
