package grpc

import (
	pb "template/api"
	"template/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"template/internal/server"
	"context"
	"fmt"
)

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	pb.RegisterDemoServer(ws.Server(), svc)
	RegisterGRPCService("demo.service", []string{rc.Server.Addr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

func RegisterGRPCService(appID string, addrs []string) {
	if cli, err := server.DiscoveryService(); err != nil {
		panic(err)
	} else if resp, err := cli.Register(context.Background(), &pb.RegSvcReqs{
		AppID: appID,
		Urls: addrs,
	}); err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}
}