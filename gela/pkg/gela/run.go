package gela

import (
	"flag"
	"fmt"
	"time"

	"github.com/chrisdoherty4/etxea/pkg/etxea"
	"github.com/chrisdoherty4/etxea/pkg/flags"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea plugin.
func Run(args []string) error {
	var opts = struct {
		name string
	}{}

	// Build a flag set
	flagSet := flag.NewFlagSet("gela", flag.ContinueOnError)
	flagSet.StringVar(&opts.name, "name", "", "Name is the name to be printed in gela plugin")

	// Create a flag registrar and register the flag set against the desired command.
	registrar := flags.NewRegistrar()
	registrar.RegisterFlags(flags.CreateCluster, flags.FromFlagSet(flagSet), flagSet)

	// Build the plugins to be registered with the plugin.ServeConfig.
	plugins := etxea.NewPluginSet(etxea.PluginDefinition{
		Flags: flags.NewPlugin(registrar),
	})

	// Launch a goroutine to print the opts - this isn't thread safe - I don't care.
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Printf("Opts: %+v\n", opts)
	}()

	// Serve the plugin so the host proces can start using us.
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: etxea.Handshake,
		Plugins:         plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
		Logger:          hclog.Default(),
	})

	return nil
}
