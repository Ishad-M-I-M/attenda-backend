package controllers

import (
	"attenda_backend/db"
	"attenda_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateClass(c *gin.Context) {
	var class models.Class

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&class)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

func GetClasses(c *gin.Context) {
	var classes []models.Class

	result := db.DB.Find(&classes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, classes)
}
