package main

import (
	// newsd "first_project/internal"
	"first_project/internal/config"
	"fmt"

	"log/slog"
	"os"
	"database/sql"


	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
)

const (
	envLocal = "local"
	envProd = "prod"
)

var DB *sql.DB

const (
    HOST = "localhost"
    PORT = 5432
    USER = "myuser"
    PASSWORD = "mysecretpassword"
    DBNAME = "mydatabase"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("Starting URl-shortener", slog.String("env", cfg.Env))
	log.Debug("Debug mess are enabled")

	//db
	 connString := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        HOST, PORT, USER, PASSWORD, DBNAME,
    )
	DB, err := sql.Open("postgres", connString)
    if err != nil {
        log.Error(err.Error())
    }

	err = DB.Ping()
    if err != nil {
        log.Error("ping error")
    }

	rows, err := DB.Query("SELECT name FROM users")
	if err != nil {
		log.Error("Query error: %", err.Error())
	}
	defer rows.Close()

    defer DB.Close()

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