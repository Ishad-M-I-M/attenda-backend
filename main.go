package main

import (
	"attenda_backend/controllers"
	"attenda_backend/db"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	err = db.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "attenda_backend",
		})
	})

	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("/", controllers.GetStudents)
		studentRoutes.POST("/", controllers.CreateStudent)
		studentRoutes.GET("/attendance", controllers.GetStudentAttendance)
	}

	teachersRoutes := router.Group("/teachers")
	{
		teachersRoutes.GET("/", controllers.GetTeacher)
		teachersRoutes.POST("/", controllers.CreateTeacher)
	}

	classRoutes := router.Group("/classes")
	{
		classRoutes.GET("/", controllers.GetClasses)
		classRoutes.POST("/", controllers.CreateClass)
		classRoutes.GET("/attendance", controllers.GetAttendance)
	}

	studentClassRoutes := router.Group("/student_classes")
	{
		studentClassRoutes.POST("/assign", controllers.Assign)
		studentClassRoutes.POST("/attendance", controllers.MarkAttendance)
	}

	router.Run()
}
