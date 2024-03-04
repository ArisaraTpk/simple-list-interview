package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
)

type UpdateInterviewDetailSvc interface {
	Execute(req UpdateInterviewDetailReq, ctx *gin.Context, l zerolog.Logger) (*UpdateInterviewDetailRes, *errors.APIError)
}

type UpdateInterviewDetailReq struct {
	InterviewId int32   `json:"interviewId" validate:"required,numeric"`
	UserId      string  `json:"_" validate:"required"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	IsArchive   *bool   `json:"isArchive,omitempty"`
}

type UpdateInterviewDetailRes struct {
	Message string `json:"message"`
}
