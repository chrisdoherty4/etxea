package etxea

import (
	"context"
	"os/exec"

	etxeav1 "github.com/chrisdoherty4/etxea/pkg/api/etxea/v1"
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

	rawClient, err := rpcClient.Dispense(BindingsPluginName)
	if err != nil {
		return err
	}

	client := rawClient.(etxeav1.BindingServiceClient)
	if _, err := client.Bind(context.Background(), &etxeav1.BindRequest{}); err != nil {
		return err
	}

	// Launch plugin
	// Initialize bindings.
	// Invoke workflow.

	return nil
}
