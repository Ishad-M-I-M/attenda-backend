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
	assignStudentToDefaultClass(&student)
	c.JSON(http.StatusOK, student)
}

func assignStudentToDefaultClass(student *models.Student) {
	// Assign student to the default class (If a default class exists) based on the details
	var defaultClasses []models.DefaultClass
	db.DB.Find(&defaultClasses, "grade = ? AND medium = ?", student.Grade, student.Medium)

	var classId uint
	if len(defaultClasses) > 0 {
		for _, defaultClass := range defaultClasses {
			if student.Gender == defaultClass.Gender || defaultClass.Gender == "mixed" {
				classId = defaultClass.ClassId
				break
			}
		}
	}

	if classId > 0 {
		studentClass := models.StudentClass{
			StudentId: student.ID,
			ClassId:   classId,
		}
		db.DB.Create(&studentClass)
	}
}
