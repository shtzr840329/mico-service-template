package http

import (
	"net/http"

	pb "template/api"
	"template/internal/model"
	"template/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"template/internal/server"
	"template/internal/utils"
	"context"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	pb.RegisterDemoBMServer(engine, svc)
	RegisterHTTPService()
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func RegisterHTTPService() {
	if dsSvc, err := server.DiscoveryService(); err != nil {
		log.Error("Fetch discovery service error: %v", err)
	} else if data, err := utils.PickPathsFromSwaggerJSON("/Users/zhaojiachen/Projects/template/api/api.swagger.json"); err != nil {
		log.Error("API swagger file open failed: %v", err)
	} else if _, err := dsSvc.AddRoutes(context.Background(), &pb.AddRoutesReqs{
		ServiceID: "discovery.service",
		Paths: data,
	}); err != nil {
		panic(err)
	}
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/template")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

