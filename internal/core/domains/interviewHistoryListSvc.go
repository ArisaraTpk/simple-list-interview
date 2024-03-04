package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
	"time"
)

type InterviewHistoryListSvc interface {
	Execute(req InterviewHistoryListReq, ctx *gin.Context, l zerolog.Logger) (*InterviewHistoryListRes, *errors.APIError)
}

type InterviewHistoryListReq struct {
	InterviewId int32 `json:"interviewId" validate:"required,numeric"`
}

type InterviewHistoryListRes struct {
	InterviewHistoryList []InterviewHistoryList `json:"interviewList"`
}

type InterviewHistoryList struct {
	HistoryId   int32     `json:"historyId"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	UpdatedBy   string    `json:"updatedBy"`
	Status      *string   `json:"status"`
	IsArchive   *bool     `json:"isArchive"`
	CreatedAt   time.Time `json:"createdAt"`
}
