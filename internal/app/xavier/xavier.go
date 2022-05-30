package xavier

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/chrisdoherty4/xavier/internal/pkg/etxea"
	"github.com/chrisdoherty4/xavier/internal/pkg/xavier"
	xavierv1 "github.com/chrisdoherty4/xavier/pkg/api/xavier/v1"
	"github.com/hashicorp/go-plugin"
)

// Run executes the xavier program.
func Run(args []string) error {
	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  xavier.Handshake,
		Plugins:          etxea.Plugins,
		Cmd:              exec.Command(args[1]),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer pluginClient.Kill()

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return err
	}

	rawClient, err := rpcClient.Dispense(etxea.PluginName)
	if err != nil {
		return err
	}

	client := rawClient.(xavierv1.ProviderPluginServiceClient)
	response, err := client.Init(context.Background(), &xavierv1.InitRequest{Message: "client sent init request"})
	if err != nil {
		return err
	}

	fmt.Printf("client: %v\n", response)

	return nil
}
