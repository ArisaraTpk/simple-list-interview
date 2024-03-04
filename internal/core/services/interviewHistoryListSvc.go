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

type interviewHistoryListSvc struct {
	validator *validator.Validate
	history   ports.InterviewHistoryRepo
}

func NewInterviewHistoryListSvc(validator *validator.Validate, history ports.InterviewHistoryRepo) domains.InterviewHistoryListSvc {
	return &interviewHistoryListSvc{
		validator: validator,
		history:   history,
	}
}

func (s interviewHistoryListSvc) Execute(req domains.InterviewHistoryListReq, ctx *gin.Context, l zerolog.Logger) (*domains.InterviewHistoryListRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	reqInterview := ports.FindInterviewHistoryReq{
		InterviewId: req.InterviewId,
	}
	interviewHistoryList, err := s.history.FindInterviewHistory(reqInterview)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("call FindInterviewHistory failed %v", err))
		return nil, errors.ErrTechnical
	}

	return s.buildInterviewHistoryListRes(interviewHistoryList), nil
}

func (s interviewHistoryListSvc) validate(req domains.InterviewHistoryListReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s interviewHistoryListSvc) buildInterviewHistoryListRes(interviewHistoryList []ports.InterviewHistoryEntity) *domains.InterviewHistoryListRes {
	list := []domains.InterviewHistoryList{}
	for _, interview := range interviewHistoryList {
		data := domains.InterviewHistoryList{
			HistoryId:   interview.HistoryId,
			Title:       interview.Title,
			Description: interview.Description,
			UpdatedBy:   interview.UpdatedBy,
			Status:      interview.Status,
			IsArchive:   interview.IsArchive,
			CreatedAt:   interview.CreatedAt,
		}
		list = append(list, data)
	}
	return &domains.InterviewHistoryListRes{InterviewHistoryList: list}
}
