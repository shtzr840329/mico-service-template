package server

import (
	pb "template/api"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"fmt"
	"context"
)

func DiscoveryService() (pb.DiscoveryClient, error) {
	cli := warden.NewClient(nil)
	conn, err := cli.Dial(context.Background(), "discovery://default/discovery.service")
	if err != nil {
		return nil, fmt.Errorf("Register center is unready: %v\n", err)
	}
	return pb.NewDiscoveryClient(conn), nil
}