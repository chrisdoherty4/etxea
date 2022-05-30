package gela

import (
	"log"
	"strings"

	"github.com/chrisdoherty4/etxea/pkg/etxea"
	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea plugin.
func Run(args []string) error {

	plugins := plugin.PluginSet{
		etxea.FlagsPluginName:    &etxea.FlagsPlugin{},
		etxea.BindingsPluginName: &etxea.BindingsPlugin{},
	}

	var names []string
	for plugin := range plugins {
		names = append(names, plugin)
	}
	log.Printf("Supported plugins: %v\n", strings.Join(names, ", "))

	plugin.Serve(NewServeConfig(plugins))

	return nil
}
