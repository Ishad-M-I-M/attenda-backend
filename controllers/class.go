package controllers

import (
	"attenda_backend/db"
	"attenda_backend/dtos"
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

	db.DB.Preload("Teacher").First(&class, class.ID)

	c.JSON(http.StatusOK, class)
}

func GetClasses(c *gin.Context) {
	var classes []models.Class

	result := db.DB.Preload("Teacher").Find(&classes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetAttendance(c *gin.Context) {
	var attendance []models.Attendance

	classId := c.Query("class_id")
	date := c.Query("date")
	if classId == "" || date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "class_id and date are required"})
		return
	}

	result := db.DB.Where("class_id = ? AND date = ?", classId, date).
		Preload("Class").
		Find(&attendance)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	var studentclasses []models.StudentClass
	result = db.DB.Preload("Student").Find(&studentclasses, "class_id = ?", classId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	classAttendance := dtos.ClassAttendance{}
	classAttendance.ClassId = attendance[0].ClassId
	classAttendance.ClassName = attendance[0].Class.Name
	classAttendance.Date = attendance[0].Date

	attendedStudentIds := map[uint]struct{}{}
	for _, a := range attendance {
		attendedStudentIds[a.StudentId] = struct{}{}
	}

	for _, sc := range studentclasses {
		present := false
		if _, exists := attendedStudentIds[sc.Student.ID]; exists {
			present = true
		}

		classAttendance.Students = append(classAttendance.Students, dtos.Student{
			StudentId:   sc.Student.ID,
			StudentName: sc.Student.Name,
			Present:     present,
		})
	}

	c.JSON(http.StatusOK, classAttendance)
}
