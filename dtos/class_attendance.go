package dtos

import "time"

type Student struct {
	StudentId   uint
	StudentName string
	Present     bool
}

type ClassAttendance struct {
	ClassId   uint      `json:"class_id"`
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"`
	Date      time.Time `json:"date"`
}
