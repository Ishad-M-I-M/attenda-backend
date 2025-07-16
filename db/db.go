package db

import (
	"attenda_backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Enable foreign key constraints
	DB.Exec("PRAGMA foreign_keys = ON")
	err = DB.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Class{}, &models.StudentClass{}, &models.DefaultClass{})
	if err != nil {
		return err
	}

	return nil
}
