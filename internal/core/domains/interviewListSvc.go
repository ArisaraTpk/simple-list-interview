package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
	"time"
)

type InterviewListSvc interface {
	Execute(req InterviewListReq, ctx *gin.Context, l zerolog.Logger) (*InterviewListRes, *errors.APIError)
}

type InterviewListReq struct {
	LastItemOrder int `json:"lastItemOrder" validate:"required,numeric"`
	Size          int `json:"size" validate:"required,numeric"`
	UserId        string
}

type InterviewListRes struct {
	InterviewList []InterviewList `json:"interviewList"`
}

type InterviewList struct {
	InterviewId     int32     `json:"interviewId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CreatedBy       string    `json:"createdBy"`
	Status          string    `json:"status"`
	AppointmentDate time.Time `json:"appointmentDate"`
}
