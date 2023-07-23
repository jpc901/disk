package main

import (
	"account/handler"
	account "account/proto"
	"fmt"
	"net"

	"github.com/jpc901/disk-common/conf"
	"github.com/jpc901/disk-common/db"
	"google.golang.org/grpc"
)


func Init() {
	conf.GetConfig().InitConfig(".")
	db.GetDBInstance().Init(*conf.GetConfig().MySQLConfig)
}

func main() {
	Init()
	lister, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GetConfig().AccountConfig.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	rpcServer := grpc.NewServer()
	account.RegisterUserServiceServer(rpcServer, &handler.User{})
	// 启动服务
	err = rpcServer.Serve(lister)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}