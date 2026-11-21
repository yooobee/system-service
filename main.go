package main

import (
	"flag"
	"fmt"
	configserviceServer "system-service/internal/server/configservice"

	"system-service/internal/config"
	deptserver "system-service/internal/server/deptservice"
	menuserver "system-service/internal/server/menuservice"
	roleserver "system-service/internal/server/roleservice"
	userserver "system-service/internal/server/userservice"
	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 一个 system 微服务，注册四个业务服务
		system.RegisterUserServiceServer(grpcServer, userserver.NewUserServiceServer(ctx))
		system.RegisterRoleServiceServer(grpcServer, roleserver.NewRoleServiceServer(ctx))
		system.RegisterDeptServiceServer(grpcServer, deptserver.NewDeptServiceServer(ctx))
		system.RegisterMenuServiceServer(grpcServer, menuserver.NewMenuServiceServer(ctx))
		system.RegisterConfigServiceServer(grpcServer, configserviceServer.NewConfigServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
