package dtos

type AttendanceSummaryResponse struct {
	Total   int64          `json:"total"`
	Present int64          `json:"present"`
	Date    string         `json:"date"`
	Classes []ClassSummary `json:"classes"`
}

type ClassSummary struct {
	Class   string `json:"class"`
	Total   int64  `json:"total"`
	Present int64  `json:"present"`
}
