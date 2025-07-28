package models

import (
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	gorm.Model
	StudentId uint      `json:"student_id" gorm:"not null"`
	ClassId   uint      `json:"class_id" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null;type:date;default:CURRENT_DATE"`

	Student Student `json:"student" gorm:"foreignKey:StudentId"`
	Class   Class   `json:"class" gorm:"foreignKey:ClassId"`
}
