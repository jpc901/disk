package main

import (
	"context"
	"disk-master/global"
	"disk-master/router"
	"disk-master/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	global.Config.InitConfig(".")
	log.Debug(global.Config.ServerConfig.Port)
	global.DB.Init(*global.Config.MySQLConfig)
	global.RDB.Init(*global.Config.RedisConfig)
	defer global.RDB.GetConn().Close()
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