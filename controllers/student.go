package controllers

import (
	"attenda_backend/db"
	"attenda_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	var d = db.DB
	if name := c.Query("name"); name != "" {
		d = d.Where("name LIKE ?", "%"+name+"%")
	}
	if grade := c.Query("grade"); grade != "" {
		d = d.Where("grade = ?", grade)
	}
	if medium := c.Query("medium"); medium != "" {
		d = d.Where("medium = ?", medium)
	}
	if gender := c.Query("gender"); gender != "" {
		d = d.Where("gender = ?", gender)
	}

	result := d.Find(&students)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}
