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

	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//
	//var level zapcore.Level
	//if err := level.UnmarshalText([]byte(cfg.AppConfig.LoggerConfig.Level)); err != nil {
	//	fmt.Fprintf(os.Stderr, "failed to parse log level: %v", err)
	//	fmt.Fprintf(os.Stderr, "setting log level to error")
	//	level, _ = zapcore.ParseLevel("error")
	//}
	//
	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(encoderConfig),
	//	zapcore.Lock(os.Stdout),
	//	level,
	//)
	//var logger *zap.logger
	//
	//logger = zap.New(core)
	//defer logger.Sync()

	logger := adapters.NewZapLogger(cfg.AppConfig.LoggerConfig)
	//logger, _ := zap.NewProduction()
	logger.Info("test")
	defer logger.Sync()
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	//	Output:    logger.WithOptions(zap.AddCallerSkip(1)).Writer(),
	//	SkipPaths: []string{"/ping"},
	//}))
	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	})
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
