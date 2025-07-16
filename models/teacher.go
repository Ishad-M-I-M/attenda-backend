package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Gender Gender `json:"gender" gorm:"not null"`
	Mobile string `json:"mobile" gorm:"not null"`
}

func (Teacher) TableName() string {
	return "teachers"
}
