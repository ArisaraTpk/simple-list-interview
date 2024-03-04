package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	"simple-list-interview/internal/core/domains"
	"simple-list-interview/internal/core/ports"
	"simple-list-interview/middleware/errors"
	"time"
)

type updateInterviewDetailSvc struct {
	validator        *validator.Validate
	interviews       ports.InterviewsRepo
	interviewHistory ports.InterviewHistoryRepo
}

func NewUpdateInterviewDetailSvc(validator *validator.Validate, interviews ports.InterviewsRepo, interviewHistory ports.InterviewHistoryRepo) domains.UpdateInterviewDetailSvc {
	return &updateInterviewDetailSvc{
		validator:        validator,
		interviews:       interviews,
		interviewHistory: interviewHistory,
	}
}

func (s updateInterviewDetailSvc) Execute(req domains.UpdateInterviewDetailReq, ctx *gin.Context, l zerolog.Logger) (*domains.UpdateInterviewDetailRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}
	oldInterviewDetail, err := s.FindOldInterviewDetail(req, l)
	if err != nil {
		return nil, err
	}

	errUpdate := s.UpdateInterviewDetail(req, l)
	if errUpdate != nil {
		return nil, errUpdate
	}
	errCreate := s.createEditHistory(req, oldInterviewDetail, l)
	if errCreate != nil {
		return nil, errCreate
	}

	return s.buildInterviewDetailRes(), nil
}

func (s updateInterviewDetailSvc) validate(req domains.UpdateInterviewDetailReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s updateInterviewDetailSvc) buildInterviewDetailRes() *domains.UpdateInterviewDetailRes {
	return &domains.UpdateInterviewDetailRes{
		Message: "Success",
	}
}

func (s updateInterviewDetailSvc) createEditHistory(newData domains.UpdateInterviewDetailReq, oldData *ports.InterviewsEntity, l zerolog.Logger) *errors.APIError {
	editHistoryData := ports.InterviewHistoryEntity{
		InterviewId: oldData.InterviewId,
		UpdatedBy:   newData.UserId,
		CreatedAt:   time.Now(),
	}
	if newData.Title != nil {
		editHistoryData.Title = &oldData.Title
	}
	if newData.Description != nil {
		editHistoryData.Description = &oldData.Description
	}
	if newData.Status != nil {
		editHistoryData.Status = &oldData.Status
	}
	if newData.IsArchive != nil {
		editHistoryData.IsArchive = &oldData.IsArchive
	}

	err := s.interviewHistory.CreateInterviewHistory(editHistoryData)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("createEditHistory errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s updateInterviewDetailSvc) UpdateInterviewDetail(req domains.UpdateInterviewDetailReq, l zerolog.Logger) *errors.APIError {
	reqUpdate := ports.UpdateInterviewDetailReq{
		InterviewId: req.InterviewId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		IsArchive:   req.IsArchive,
		UpdatedBy:   req.UserId,
	}
	err := s.interviews.UpdateInterviewDetail(reqUpdate)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("UpdateInterviewDetail errors %v", err))
		return errors.ErrTechnical
	}
	return nil
}

func (s updateInterviewDetailSvc) FindOldInterviewDetail(req domains.UpdateInterviewDetailReq, l zerolog.Logger) (*ports.InterviewsEntity, *errors.APIError) {
	reqOldData := ports.FindInterviewDetailReq{InterviewId: req.InterviewId}
	oldData, err := s.interviews.FindInterviewDetail(reqOldData)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("FindInterviewDetail errors %v", err))
		return nil, errors.ErrTechnical
	}
	return oldData, nil
}
