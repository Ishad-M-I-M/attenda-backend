package models

import "time"

type StudentClass struct {
	StudentId  uint      `json:"student_id" gorm:"not null"`
	ClassId    uint      `json:"class_id" gorm:"not null"`
	EnrolledAt time.Time `json:"enrolled_at" gorm:"default:CURRENT_TIMESTAMP"`

	Student Student `gorm:"foreignKey:StudentId"`
	Class   Class   `gorm:"foreignKey:ClassId"`
}
