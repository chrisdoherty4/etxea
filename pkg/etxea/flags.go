package etxea

import (
	"context"

	etxeav1 "github.com/chrisdoherty4/etxea/pkg/api/etxea/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

const FlagsPluginName = "flags"

type FlagsPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

func (f *FlagsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	etxeav1.RegisterFlagServiceServer(s, &FlagsService{})
	return nil
}

func (f *FlagsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return etxeav1.NewBindingServiceClient(c), nil
}

// FlagsService is the actual implemetation.
type FlagsService struct{}

func (f *FlagsService) Define(context.Context, *etxeav1.Empty) (*etxeav1.Empty, error) {
	return &etxeav1.Empty{}, nil
}
