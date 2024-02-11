package main

import (
	"healthcare-capt-america/apps"
	"healthcare-capt-america/middlewares"
	"healthcare-capt-america/pkg/configs"
	"healthcare-capt-america/pkg/databases"
	"healthcare-capt-america/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	envFilePath := "./files/env/.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		return
	}
}

func main() {
	configPath, err := configs.ParseFlags()
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.LoggerMiddleware())
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("cokkk", gin.Mode())
	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	return
	}
	repo, err := databases.NewRepositories(config)
	if err != nil {
		return
	}
	repo.InitRepositories()
	cron := services.NewCronService(repo.OrderRepository)
	err = cron.RunAllJob()
	if err != nil {
		log.Fatal(err)
		return
	}
	cron.Scheduler.Start()
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})
	services.InitAuthority(repo.GetDB())
	handlers := apps.NewHandlers(repo)
	handlers.InitRouter(r)
	config.Run(r)
}
