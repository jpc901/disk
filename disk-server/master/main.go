package main

import (
	"context"
	"master/global"
	"master/grpc/account"
	ossClient "master/oss"
	"master/router"
	"master/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func Init() {
	global.Config.InitConfig(".")
	log.Debug(global.Config.ServerConfig.Port)
	global.DB.Init(*global.Config.MySQLConfig)
	global.RDB.Init(*global.Config.RedisConfig)
	err := account.InitClient()
	if err != nil {
		log.Errorf("account init failed %v", err)
	}
	ossClient.GetClient().Init()
}

func main() {
	Init()
	defer global.RDB.GetConn().Close()

	go ossClient.GetClient().UploadFile()

	// start http server
	router := router.Init()
	httpServer := server.GetHttpServerInstance(router)
	httpServer.Start()

	// exit
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Merak Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.ShutDown(ctx)
}
