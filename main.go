package main

import (
	// newsd "first_project/internal"
	"first_project/internal/config"
	"fmt"

	"log/slog"
	"os"


	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("Starting URl-shortener", slog.String("env", cfg.Env))
	log.Debug("Debug mess are enabled")

	//db

	router := gin.Default()
	router.GET("/", func(c *gin.Context) { 
		c.JSON(200, 
			gin.H{ "message": "hello world",  
			})  
		})
	router.POST("/upload", func(c *gin.Context) {
		
	})
	router.Run(cfg.Address)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}