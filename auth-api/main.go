package main

import (
	"auth-api/app/config"
	"auth-api/app/container"
	"auth-api/app/http/routes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pickme-go/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {

	cfg := config.Parse("./config")
	ctr := container.Resolve(cfg)

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	log.Info("Initializing routes")
	routes.InitRoutes(router, ctr)

	log.Info("Starting server on port: " + strconv.Itoa(cfg.AppConfig.Port))

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.AppConfig.Port),
		Handler: router,
	}
	server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wait time.Duration
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)

	os.Exit(0)
}
