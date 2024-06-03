package main

import (
	"auth-api/app/config"
	"auth-api/app/container"
	"auth-api/app/http/routes"
	"auth-api/external/adapters"
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

	log := adapters.NewZapLogger(cfg.AppConfig.LoggerConfig)
	defer log.Sync()

	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Set("logger", log)
		c.Next()
	})
	routes.InitRoutes(router, ctr)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.AppConfig.Port),
		Handler: router,
	}
	log.Info("Starting server on port " + strconv.Itoa(cfg.AppConfig.Port))
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
