package main

import (
	"alumni-pendaftaran/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{"Success": true})
		})
	}

	r.Run()
}
