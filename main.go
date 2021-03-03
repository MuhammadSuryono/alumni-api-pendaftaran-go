package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{"Success": true})
		})
	}
	
	r.Run()
}
