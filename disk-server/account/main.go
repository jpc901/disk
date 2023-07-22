package main

import (
	"fmt"
	"net"

	// "disk-server/account/handler"
	"disk-server/account/handler"
	account "disk-server/account/proto"

	"github.com/jpc901/disk-common/conf"
	"github.com/jpc901/disk-common/db"
	"google.golang.org/grpc"
)


func Init() {
	conf.GetConfig().InitConfig("../config.yaml")
	db.GetDBInstance().Init(*conf.GetConfig().MySQLConfig)
}

func main() {
	lister, err := net.Listen("tcp", ":8972")
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