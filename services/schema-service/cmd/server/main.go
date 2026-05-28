package main

import (
	"schema-service/internal/db"
	"schema-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/schema", handler.CreateSchema)

	r.Run(":8081")
}