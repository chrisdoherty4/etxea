package etxea

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

// Run executes the etxea program.
func Run(args []string) error {
	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  Handshake,
		Plugins:          Plugins,
		Cmd:              exec.Command(args[1]),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer pluginClient.Kill()

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return err
	}

	_, err = rpcClient.Dispense(BindingsPluginName)
	if err != nil {
		return err
	}

	// Launch plugin
	// Initialize bindings.
	// Invoke workflow.

	return nil
}
