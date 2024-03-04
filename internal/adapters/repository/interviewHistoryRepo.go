package repository

import (
	"gorm.io/gorm"
	"simple-list-interview/internal/core/ports"
)

type interviewHistoryRepo struct {
	db *gorm.DB
}

func NewInterviewHistoryRepo(db *gorm.DB) ports.InterviewHistoryRepo {
	return &interviewHistoryRepo{
		db: db,
	}
}

func (r interviewHistoryRepo) FindInterviewHistory(req ports.FindInterviewHistoryReq) ([]ports.InterviewHistoryEntity, error) {
	var result []ports.InterviewHistoryEntity
	res := r.db.Exec(`SELECT * FROM interviewHistory WHERE interviewId = ? ORDER BY createdAt DESC`,
		req.InterviewId).Find(&result)
	return result, res.Error
}

func (r interviewHistoryRepo) CreateInterviewHistory(data ports.InterviewHistoryEntity) error {
	res := r.db.Table("interviewHistory").Create(&data)
	return res.Error
}
