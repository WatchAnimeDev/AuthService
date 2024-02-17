package main

import (
	"auth-service/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	router := gin.Default()
	router.GET("/uid/set", routers.SetUID)
	router.GET("/uid/map", routers.MapUID)
	router.Run()
}
