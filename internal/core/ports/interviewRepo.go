package ports

import "time"

type InterviewsRepo interface {
	FindActiveInterviews(req FindActiveInterviewsReq) ([]InterviewsEntity, error)
	FindInterviewDetail(req FindInterviewDetailReq) (*InterviewsEntity, error)
	UpdateInterviewDetail(req UpdateInterviewDetailReq) error
}

type FindActiveInterviewsReq struct {
	UserId    string
	LastOrder int
	Size      int
}

type FindInterviewDetailReq struct {
	InterviewId int32
}

type UpdateInterviewDetailReq struct {
	InterviewId int32
	Title       *string
	Description *string
	Status      *string
	IsArchive   *bool
	UpdatedBy   string
	UpdatedAt   time.Time
}

type InterviewsEntity struct {
	InterviewId     int32     `gorm:"column:interviewId;primaryKey;autoIncrement:true"`
	Title           string    `gorm:"column:title"`
	Description     string    `gorm:"column:description"`
	CreatedBy       string    `gorm:"column:createdBy"`
	UpdatedBy       string    `gorm:"column:updatedBy"`
	Status          string    `gorm:"column:status"`
	IsArchive       bool      `gorm:"column:isArchive"`
	AppointmentDate time.Time `gorm:"column:appointmentDate"`
	CreatedAt       time.Time `gorm:"column:createdAt"`
	UpdatedAt       time.Time `gorm:"column:updatedAt"`
}
