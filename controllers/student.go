package controllers

import (
	"attenda_backend/db"
	"attenda_backend/dtos"
	"attenda_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetStudentAttendance(c *gin.Context) {
	studentId := c.Query("student_id")
	classId := c.Query("class_id")
	limit := c.Query("limit")
	if studentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	d := db.DB.Where("student_id = ?", studentId)

	if classId != "" {
		d = d.Where("class_id = ?", classId)
	}

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a positive integer"})
			return
		}
		d = d.Limit(limitInt)
	}

	var attendance []models.Attendance
	result := d.Order("date DESC").
		Preload("Student").
		Preload("Class").
		Find(&attendance)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	if len(attendance) == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	studentAttendace := dtos.StudentAttendance{}
	studentAttendace.StudentId = attendance[0].StudentId
	studentAttendace.StudentName = attendance[0].Student.Name

	for _, att := range attendance {
		studentAttendace.Records = append(studentAttendace.Records, dtos.AttendanceRecord{
			ClassId:   att.ClassId,
			ClassName: att.Class.Name,
			Date:      dtos.DateOnlyFromTime(att.Date),
		})
	}

	c.JSON(http.StatusOK, studentAttendace)
}
