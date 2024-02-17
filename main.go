package main

import (
	"auth-service/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.GET("/uid/set", routers.SetUID)
	router.GET("/uid/map", routers.MapUID)
	router.Run()
}
