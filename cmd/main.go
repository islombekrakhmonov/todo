package main

import (
	"fmt"
	"todo/api"
	"todo/api/handlers"
	"todo/config"
	"todo/storage/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
 cfg := config.Load()

 pgStore := postgres.NewPostgres(fmt.Sprintf(
  "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
  cfg.PostgresHost,
  cfg.PostgresPort,
  cfg.PostgresUser,
  cfg.PostgresPassword,
  cfg.PostgresDatabase,
 ))

 h := handlers.NewHandler(pgStore, cfg)

 switch cfg.Environment { // developement, test, production
 case "dev":
  gin.SetMode(gin.DebugMode)
 case "test":
  gin.SetMode(gin.TestMode)
 default:
  gin.SetMode(gin.ReleaseMode)
 }

 r := gin.New()

 r.Use(gin.Logger(), gin.Recovery())

 api.SetUpAPI(r, h, cfg)

 r.Run(cfg.HTTPPort)
}