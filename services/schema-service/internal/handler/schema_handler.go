package handler

import (
	"net/http"

	"../model"
	"../service"

	"github.com/gin-gonic/gin"
)

func CreateSchema(c *gin.Context) {
	var input model.Schema

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateSchema(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schema created"})
}