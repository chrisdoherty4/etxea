package xavier

import (
	"context"

	xavierv1 "github.com/chrisdoherty4/xavier/internal/pkg/api/xavier/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

const BindingsPluginName = "bindings"

type BindingsPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

func (p *BindingsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	xavierv1.RegisterBindingServiceServer(s, &BindingsService{})
	return nil
}

func (p *BindingsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return xavierv1.NewBindingServiceClient(c), nil
}

type BindingsService struct{}
