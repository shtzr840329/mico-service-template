package server

import (
	"context"
	"fmt"
	pb "template/api"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

func RegisterService() (pb.RegisterClient, error) {
	cli := warden.NewClient(nil)
	conn, err := cli.Dial(context.Background(), "discovery://default/register.service")
	if err != nil {
		return nil, fmt.Errorf("Register center is unready: %v\n", err)
	}
	return pb.NewRegisterClient(conn), nil
}
