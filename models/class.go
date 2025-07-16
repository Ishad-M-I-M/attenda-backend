package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	TeacherId   uint    `json:"teacher_id"`
	Teacher     Teacher `json:"teacher" gorm:"foreignKey:TeacherId"`

	StudentClasses []StudentClass `gorm:"foreignKey:ClassId"`
}

func (Class) TableName() string {
	return "classes"
}

// DefaultClass represents a class with a specific grade, medium or gender if seperated by.
// Data will be seeded initially
// Used to assign students at the time of registration.
type DefaultClass struct {
	gorm.Model
	Grade   uint8  `json:"grade" gorm:"not null"`
	Medium  Medium `json:"medium" gorm:"not null"`
	Gender  Gender `json:"gender"`
	ClassId uint   `json:"class_id" gorm:"not null"`
	Class   Class  `json:"class" gorm:"foreignKey:ClassId"`
}

func (DefaultClass) TableName() string {
	return "default_classes"
}
