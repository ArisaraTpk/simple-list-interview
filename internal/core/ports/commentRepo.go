package ports

import "time"

type CommentRepo interface {
	FindCommentList(req FindCommentListReq) ([]CommentEntity, error)
	UpdateCommentDescription(req UpdateCommentDescriptionReq) error
	FindCommentDetail(req FindCommentDetailReq) (*CommentEntity, error)
}

type FindCommentListReq struct {
	InterviewId int32
}

type UpdateCommentDescriptionReq struct {
	CommentId   int32
	Description string
}

type FindCommentDetailReq struct {
	CommentId int32
}

type CommentEntity struct {
	CommentId   int32     `gorm:"column:commentId;primaryKey;autoIncrement:true"`
	InterviewId int32     `gorm:"column:interviewId"`
	CreatedBy   string    `gorm:"column:createdBy"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}
