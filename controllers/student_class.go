package controllers

import (
	"attenda_backend/db"
	"attenda_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Assign(c *gin.Context) {
	var studentClass models.StudentClass

	if err := c.ShouldBindJSON(&studentClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&studentClass)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, studentClass)
}
