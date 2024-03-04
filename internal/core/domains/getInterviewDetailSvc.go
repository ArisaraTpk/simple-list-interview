package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/internal/core/ports"
	"simple-list-interview/middleware/errors"
	"time"
)

type GetInterviewDetailSvc interface {
	Execute(req GetInterviewDetailReq, ctx *gin.Context, l zerolog.Logger) (*GetInterviewDetailRes, *errors.APIError)
}

type GetInterviewDetailReq struct {
	InterviewId int32  `json:"interviewId" validate:"required,numeric"`
	UserId      string `json:"_" validate:"required"`
}

type GetInterviewDetailRes struct {
	InterviewId     int32     `json:"interviewId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CreatedBy       string    `json:"createdBy"`
	CreatedEmail    string    `json:"createdEmail"`
	Status          string    `json:"status"`
	AppointmentDate time.Time `json:"appointmentDate"`
}

type GetInterviewDetailAll struct {
	GetInterviewDetail *ports.InterviewsEntity
	GetUserDetail      *ports.UserEntity
	InterviewDetailErr *errors.APIError
	UserErr            *errors.APIError
}
