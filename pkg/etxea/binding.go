package etxea

import (
	"context"

	etxeav1 "github.com/chrisdoherty4/etxea/pkg/api/etxea/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

const BindingsPluginName = "bindings"

type BindingsPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

func (p *BindingsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	etxeav1.RegisterBindingServiceServer(s, &BindingsService{})
	return nil
}

func (p *BindingsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return etxeav1.NewBindingServiceClient(c), nil
}

type BindingsService struct{}

func (s *BindingsService) Bind(context.Context, *etxeav1.Empty) (*etxeav1.Empty, error) {
	return &etxeav1.Empty{}, nil
}
