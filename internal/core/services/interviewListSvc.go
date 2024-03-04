package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	"simple-list-interview/internal/core/domains"
	"simple-list-interview/internal/core/ports"
	"simple-list-interview/middleware/errors"
)

type interviewListSvc struct {
	validator  *validator.Validate
	interviews ports.InterviewsRepo
}

func NewInterviewListSvc(validator *validator.Validate, interviews ports.InterviewsRepo) domains.InterviewListSvc {
	return &interviewListSvc{
		validator:  validator,
		interviews: interviews,
	}
}

func (s interviewListSvc) Execute(req domains.InterviewListReq, ctx *gin.Context, l zerolog.Logger) (*domains.InterviewListRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	reqInterview := ports.FindActiveInterviewsReq{
		UserId:    req.UserId,
		LastOrder: req.LastItemOrder,
		Size:      req.LastItemOrder,
	}
	interviewList, err := s.interviews.FindActiveInterviews(reqInterview)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("call FindActiveInterviews failed %v", err))
		return nil, errors.ErrTechnical
	}

	return s.buildInterviewListRes(interviewList), nil
}

func (s interviewListSvc) validate(req domains.InterviewListReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s interviewListSvc) buildInterviewListRes(interviewList []ports.InterviewsEntity) *domains.InterviewListRes {
	list := []domains.InterviewList{}
	for _, interview := range interviewList {
		data := domains.InterviewList{
			InterviewId:     interview.InterviewId,
			Title:           interview.Title,
			Description:     interview.Description,
			CreatedBy:       interview.CreatedBy,
			Status:          interview.Status,
			AppointmentDate: interview.AppointmentDate,
		}
		list = append(list, data)
	}
	return &domains.InterviewListRes{InterviewList: list}
}
