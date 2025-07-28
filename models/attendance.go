package models

import (
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	gorm.Model
	StudentId uint      `json:"student_id" gorm:"not null;uniqueIndex:idx_student_class_date"`
	ClassId   uint      `json:"class_id" gorm:"not null;uniqueIndex:idx_student_class_date"`
	Date      time.Time `json:"date" gorm:"not null;type:date;default:CURRENT_DATE;uniqueIndex:idx_student_class_date"`

	Student Student `json:"student" gorm:"foreignKey:StudentId"`
	Class   Class   `json:"class" gorm:"foreignKey:ClassId"`
}
