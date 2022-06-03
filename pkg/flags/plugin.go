package flags

import (
	"context"

	etxeav1 "github.com/chrisdoherty4/etxea/pkg/api/etxea/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const PluginName = "flags"

type flagsPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	registrar *Registrar
}

// NewPlugin is used to create a FlagsPlugin instance for plugin authors.
func NewPlugin(registrar *Registrar) plugin.Plugin {
	return &flagsPlugin{registrar: registrar}
}

func (f *flagsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	etxeav1.RegisterFlagServiceServer(s, &flagsService{registrar: f.registrar})
	return nil
}

func (f *flagsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return etxeav1.NewBindingServiceClient(c), nil
}

// flagsService is the actual implemetation.
type flagsService struct {
	registrar *Registrar
}

func (f *flagsService) GetFlags(_ context.Context, request *etxeav1.GetFlagsRequest) (*etxeav1.GetFlagsResponse, error) {
	if !f.registrar.hasFlags(CID(request.CommandId)) {
		return nil, grpc.Errorf(codes.Unavailable, "no flags for command")
	}

	f.registrar.GetFlags

	return &etxeav1.GetFlagsResponse{}, nil
}

func (f *flagsService) Parse(_ context.Context, request *etxeav1.ParseRequest) (*etxeav1.ParseResponse, error) {
	if f.registrar.hasFlags(CID(request.CommandId)) {
		return &etxeav1.ParseResponse{}, f.registrar.parse(CID(request.CommandId), request.Args)
	}

	return &etxeav1.ParseResponse{}, nil
}
