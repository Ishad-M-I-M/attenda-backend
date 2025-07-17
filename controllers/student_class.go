package controllers

import (
	"attenda_backend/db"
	"attenda_backend/dtos"
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

func MarkAttendance(c *gin.Context) {
	var markAttendance dtos.MarkAttendance

	if err := c.ShouldBindJSON(&markAttendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var attendance []models.Attendance
	for _, studentId := range markAttendance.StudentIds {
		attendance = append(attendance, models.Attendance{
			StudentId: studentId,
			ClassId:   markAttendance.ClassId,
		})
	}

	result := db.DB.Create(&attendance)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}
