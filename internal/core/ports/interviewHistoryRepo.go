package ports

import "time"

type InterviewHistoryRepo interface {
	FindInterviewHistory(req FindInterviewHistoryReq) ([]InterviewHistoryEntity, error)
	CreateInterviewHistory(data InterviewHistoryEntity) error
}

type FindInterviewHistoryReq struct {
	InterviewId int32
}

type InterviewHistoryEntity struct {
	HistoryId   int32     `gorm:"column:historyId;primaryKey;autoIncrement:true"`
	InterviewId int32     `gorm:"column:interviewId"`
	Title       *string   `gorm:"column:title"`
	Description *string   `gorm:"column:description"`
	UpdatedBy   string    `gorm:"column:updatedBy"`
	Status      *string   `gorm:"column:status"`
	IsArchive   *bool     `gorm:"column:isArchive"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
}
