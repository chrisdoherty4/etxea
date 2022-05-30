package etxea

import (
	"context"
	"log"

	xavierv1 "github.com/chrisdoherty4/xavier/pkg/api/xavier/v1"
	"github.com/chrisdoherty4/xavier/pkg/plugin"
	hashicorpplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var PluginName = "etxea"

var Plugins = hashicorpplugin.PluginSet{
	PluginName: &Plugin{},
}

type Plugin struct {
	hashicorpplugin.NetRPCUnsupportedPlugin
	provider plugin.Provider
}

func (p *Plugin) GRPCServer(broker *hashicorpplugin.GRPCBroker, s *grpc.Server) error {
	xavierv1.RegisterProviderPluginServiceServer(s, &Server{})
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *hashicorpplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return xavierv1.NewProviderPluginServiceClient(c), nil
}

type Server struct{}

func (c *Server) Init(ctx context.Context, in *xavierv1.InitRequest) (*xavierv1.InitResponse, error) {
	log.Printf("Init handler: %v\n", in.Message)
	return &xavierv1.InitResponse{Message: "server received init request"}, nil
}

// type Provider struct{}

// func (pp *Provider) Init(agent plugin.BindingAgent) error {
// 	var err error

// 	for stage, handler := range map[plugin.Stage]plugin.StageHandler{
// 		plugin.PreBootstrap: plugin.StageHandlerFunc(func() error {
// 			fmt.Println("PreBootstrap handler invoked")
// 			return nil
// 		}),
// 		plugin.Bootstrap: plugin.StageHandlerFunc(func() error {
// 			fmt.Println("Bootstrap handler invoked")
// 			return nil
// 		}),
// 		plugin.PostBootstrap: plugin.StageHandlerFunc(func() error {
// 			fmt.Println("PostBootstrap handler invoked")
// 			return nil
// 		}),
// 	} {
// 		err = multierr.Combine(err, agent.Bind(stage, handler))
// 	}

// 	return err
// }
