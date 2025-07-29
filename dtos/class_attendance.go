package dtos

import "time"

type Student struct {
	StudentId   uint   `json:"student_id"`
	StudentName string `json:"student_name"`
	Present     bool   `json:"present"`
}

type ClassAttendance struct {
	ClassId   uint      `json:"class_id"`
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"`
	Date      time.Time `json:"date"`
}
