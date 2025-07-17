package dtos

import "time"

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	// Remove quotes
	s = s[1 : len(s)-1]
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type MarkAttendance struct {
	Date       DateOnly `json:"date"`
	ClassId    uint     `json:"class_id"`
	StudentIds []uint   `json:"student_ids"`
}
