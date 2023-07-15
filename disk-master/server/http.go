package server

import (
	"context"
	"disk-master/global"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	port   string
	server *http.Server
}

var (
	httpServer *HttpServer
	once       sync.Once
)

func GetHttpServerInstance(Router *gin.Engine) *HttpServer {
	if httpServer != nil {
		return httpServer
	}
	port := global.Config.Port
	once.Do(func() {
		httpServer = &HttpServer{
			port: port,
			server: &http.Server{
				Addr:           fmt.Sprintf(":%s", port),
				Handler:        Router,
			},
		}
	})
	return httpServer
}

func (h *HttpServer) Start() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	log.Infof("Http Server is working on port: %s ～～～", h.port)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Http Driver listen: %s\n", err)
		}
	}()
}

func (h *HttpServer) ShutDown(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Fatal("Http Driver Shutdown:", err)
	}
}
