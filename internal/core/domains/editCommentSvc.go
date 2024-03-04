package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
)

type EditCommentSvc interface {
	Execute(req EditCommentReq, ctx *gin.Context, l zerolog.Logger) (*EditCommentRes, *errors.APIError)
}

type EditCommentReq struct {
	CommentId   int32  `json:"commentId" validate:"required,numeric"`
	Description string `json:"description" validate:"required"`
	UserId      string `json:"_" validate:"required"`
}

type EditCommentRes struct {
	Message string `json:"message"`
}
