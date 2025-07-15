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
	}

	router.Run()
}
