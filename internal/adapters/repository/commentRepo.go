package repository

import (
	"gorm.io/gorm"
	"simple-list-interview/internal/core/ports"
	"time"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) ports.CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r commentRepo) FindCommentList(req ports.FindCommentListReq) ([]ports.CommentEntity, error) {
	var result []ports.CommentEntity
	res := r.db.Exec(`SELECT * FROM comment WHERE interviewId = ? ORDER BY createdAt DESC `, req.InterviewId).Find(&result)
	return result, res.Error
}

func (r commentRepo) UpdateCommentDescription(req ports.UpdateCommentDescriptionReq) error {

	res := r.db.Exec(`UPDATE comment WHERE commentId = ? SET description = ? , updatedAt = ?`, req.CommentId, req.Description, time.Now())
	return res.Error
}

func (r commentRepo) FindCommentDetail(req ports.FindCommentDetailReq) (*ports.CommentEntity, error) {
	var result ports.CommentEntity
	res := r.db.Exec(`SELECT * FROM comment WHERE commentId = ? `, req.CommentId).First(&result)
	return &result, res.Error
}
