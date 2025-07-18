package models

import "gorm.io/gorm"

type Medium string

const (
	Sinhala Medium = "sinhala"
	Tamil   Medium = "tamil"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Mixed  Gender = "mixed"
)

type Student struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	Grade      uint8  `json:"grade" gorm:"not null"`
	Medium     Medium `json:"medium" gorm:"not null"`
	Gender     Gender `json:"gender" gorm:"not null"`
	Address    string `json:"address"`
	Mobile     string `json:"mobile" gorm:"not null"`
	FatherName string `json:"father_name"`

	StudentClasses []StudentClass `json:"-" gorm:"foreignKey:StudentId"`
}

func (Student) TableName() string {
	return "students"
}
