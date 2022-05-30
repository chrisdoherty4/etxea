package xavier

import (
	"os/exec"

	"github.com/chrisdoherty4/xavier/xavier/pkg/xavier"
	"github.com/hashicorp/go-plugin"
)

// Run executes the xavier program.
func Run(args []string) error {
	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  xavier.Handshake,
		Plugins:          xavier.Plugins,
		Cmd:              exec.Command(args[1]),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer pluginClient.Kill()

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return err
	}

	_, err = rpcClient.Dispense(xavier.BindingsPluginName)
	if err != nil {
		return err
	}

	// Launch plugin
	// Initialize bindings.
	// Invoke workflow.

	return nil
}
