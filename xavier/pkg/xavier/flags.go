package xavier

import (
	"context"

	xavierv1 "github.com/chrisdoherty4/xavier/xavier/pkg/api/xavier/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

const FlagsPluginName = "flags"

type FlagsPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

func (f *FlagsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	xavierv1.RegisterFlagServiceServer(s, &FlagsService{})
	return nil
}

func (f *FlagsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return xavierv1.NewBindingServiceClient(c), nil
}

// FlagsService is the actual implemetation.
type FlagsService struct{}
