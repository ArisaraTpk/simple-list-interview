package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
	"time"
)

type CommentListSvc interface {
	Execute(req CommentListReq, ctx *gin.Context, l zerolog.Logger) (*CommentListRes, *errors.APIError)
}

type CommentListReq struct {
	InterviewId int32  `json:"interviewId" validate:"required,numeric"`
	UserId      string `json:"_" validate:"required"`
}

type CommentListRes struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	CommentId   int32     `json:"commentId"`
	CreatedBy   string    `json:"createdBy"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
