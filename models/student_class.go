package models

import "time"

type StudentClass struct {
	StudentId  uint      `json:"student_id" gorm:"primaryKey;not null"`
	ClassId    uint      `json:"class_id" gorm:"primaryKey;not null"`
	EnrolledAt time.Time `json:"enrolled_at" gorm:"default:CURRENT_TIMESTAMP"`

	Student Student `gorm:"foreignKey:StudentId"`
	Class   Class   `gorm:"foreignKey:ClassId"`
}

func (StudentClass) TableName() string {
	return "student_classes"
}
