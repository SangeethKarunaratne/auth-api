package main

import (
	"auth-api/app/config"
	"auth-api/app/container"
	"auth-api/app/http/routes"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	cfg := config.Parse("./config")
	ctr := container.Resolve(cfg)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	routes.InitRoutes(router, ctr)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.AppConfig.Port),
		Handler: router,
	}
	server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	wait := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)

	os.Exit(0)
}
