package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/api"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/config"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/services"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

var (
	conf    *config.Config
	userSvc *services.UserSvc
	logger  *zap.Logger
)

func init() {
	// load config
	conf = config.GetConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	userSvc = services.NewUserService(conf)
}

func main() {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(r, userSvc, conf)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
