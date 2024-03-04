package repository

import (
	"fmt"
	"gorm.io/gorm"
	"simple-list-interview/internal/core/ports"
	"strings"
	"time"
)

type interviewsRepo struct {
	db *gorm.DB
}

func NewInterviewsRepo(db *gorm.DB) ports.InterviewsRepo {
	return &interviewsRepo{
		db: db,
	}
}

func (r interviewsRepo) FindActiveInterviews(req ports.FindActiveInterviewsReq) ([]ports.InterviewsEntity, error) {
	var result []ports.InterviewsEntity
	res := r.db.Exec(`SELECT * FROM interviews WHERE isArchive = false OFFSET ? ROWS FETCH NEXT ? ROWS ONLY ORDER BY createdAt DESC`,
		req.LastOrder, req.Size).Find(&result)
	return result, res.Error
}

func (r interviewsRepo) FindInterviewDetail(req ports.FindInterviewDetailReq) (*ports.InterviewsEntity, error) {
	var result ports.InterviewsEntity
	res := r.db.Exec(`SELECT * FROM interviews WHERE interviewId = ?`, req.InterviewId).Find(&result)
	return &result, res.Error
}

func (r interviewsRepo) UpdateInterviewDetail(req ports.UpdateInterviewDetailReq) error {
	sqlPatch := fmt.Sprintf(`UPDATE interviews WHERE interviewId = ? SET updatedAt = ?`, req.InterviewId, time.Now())
	qParts := []string{}
	if req.Title != nil {
		qParts = append(qParts, fmt.Sprintf(`title = %s`, *req.Title))
	}
	if req.Description != nil {
		qParts = append(qParts, fmt.Sprintf(`description = %s`, *req.Description))
	}
	if req.Status != nil {
		qParts = append(qParts, fmt.Sprintf(`status = %s`, *req.Status))
	}
	if req.IsArchive != nil {
		qParts = append(qParts, fmt.Sprintf(`isArchive = %t`, *req.IsArchive))
	}
	if req.UpdatedBy != "" {
		qParts = append(qParts, fmt.Sprintf(`UpdatedBy = %s`, req.UpdatedBy))
	}

	q := strings.Join(qParts, ",")
	sqlPatch = sqlPatch + q

	res := r.db.Exec(sqlPatch)
	return res.Error
}
