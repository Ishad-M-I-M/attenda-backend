package dtos

type AttendanceRecord struct {
	ClassId   uint     `json:"class_id"`
	ClassName string   `json:"class_name"`
	Date      DateOnly `json:"date"`
}

type StudentAttendance struct {
	StudentId   uint               `json:"student_id"`
	StudentName string             `json:"student_name"`
	Records     []AttendanceRecord `json:"records"`
}
