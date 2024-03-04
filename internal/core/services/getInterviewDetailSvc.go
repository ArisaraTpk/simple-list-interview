package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	"simple-list-interview/internal/core/domains"
	"simple-list-interview/internal/core/ports"
	"simple-list-interview/middleware/errors"
	"sync"
)

type getInterviewDetailSvc struct {
	validator  *validator.Validate
	interviews ports.InterviewsRepo
	user       ports.UserRepo
}

func NewGetInterviewDetailSvc(validator *validator.Validate, interviews ports.InterviewsRepo, user ports.UserRepo) domains.GetInterviewDetailSvc {
	return &getInterviewDetailSvc{
		validator:  validator,
		interviews: interviews,
		user:       user,
	}
}

func (s getInterviewDetailSvc) Execute(req domains.GetInterviewDetailReq, ctx *gin.Context, l zerolog.Logger) (*domains.GetInterviewDetailRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	data := s.getInterviewDetailAll(req, l)
	if data.InterviewDetailErr != nil {
		return nil, data.InterviewDetailErr
	} else if data.UserErr != nil {
		return nil, data.UserErr
	}

	return s.buildInterviewDetailRes(data.GetInterviewDetail, data.GetUserDetail), nil
}

func (s getInterviewDetailSvc) validate(req domains.GetInterviewDetailReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s getInterviewDetailSvc) buildInterviewDetailRes(interviewDetail *ports.InterviewsEntity, userDetail *ports.UserEntity) *domains.GetInterviewDetailRes {

	return &domains.GetInterviewDetailRes{
		InterviewId:     interviewDetail.InterviewId,
		Title:           interviewDetail.Title,
		Description:     interviewDetail.Description,
		CreatedBy:       interviewDetail.CreatedBy,
		CreatedEmail:    userDetail.Email,
		Status:          interviewDetail.Status,
		AppointmentDate: interviewDetail.AppointmentDate,
	}
}

func (s getInterviewDetailSvc) getInterviewDetail(req domains.GetInterviewDetailReq, l zerolog.Logger) (*ports.InterviewsEntity, *errors.APIError) {
	reqInterview := ports.FindInterviewDetailReq{
		InterviewId: req.InterviewId,
	}
	getInterviewDetail, err := s.interviews.FindInterviewDetail(reqInterview)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("call FindInterviewDetail failed %v", err))
		return nil, errors.ErrTechnical
	}

	return getInterviewDetail, nil
}

func (s getInterviewDetailSvc) getUserDetail(userId string, l zerolog.Logger) (*ports.UserEntity, *errors.APIError) {

	getUserDetail, err := s.user.FindUserDetail(userId)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("call FindUserDetail failed %v", err))
		return nil, errors.ErrTechnical
	}

	return getUserDetail, nil
}

func (s getInterviewDetailSvc) getInterviewDetailAll(req domains.GetInterviewDetailReq, l zerolog.Logger) domains.GetInterviewDetailAll {
	var wg sync.WaitGroup
	getInterviewDetail, getUserDetail := &ports.InterviewsEntity{}, &ports.UserEntity{}
	var interviewDetailErr, userErr *errors.APIError
	chInterviewDetail, chUser := make(chan *ports.InterviewsEntity), make(chan *ports.UserEntity)
	chInterviewDetailErr, chUserErr := make(chan *errors.APIError), make(chan *errors.APIError)

	wg.Add(1)
	go func(req domains.GetInterviewDetailReq, l zerolog.Logger) {
		defer wg.Done()
		getInterviewDetail, err := s.getInterviewDetail(req, l)
		if err != nil {
			l.Error().
				Err(err).
				Msg(fmt.Sprintf("get interviewDetailErr errors %v", err))
			chInterviewDetailErr <- err
			chInterviewDetail <- nil
			return
		}
		chInterviewDetailErr <- nil
		chInterviewDetail <- getInterviewDetail
	}(req, l)

	wg.Add(1)
	go func(req domains.GetInterviewDetailReq, l zerolog.Logger) {
		defer wg.Done()
		getUserDetail, err := s.getUserDetail(req.UserId, l)
		if err != nil {
			l.Error().
				Err(err).
				Msg(fmt.Sprintf("get userErr errors %v", err))
			chUserErr <- err
			chUser <- nil
			return
		}
		chUserErr <- nil
		chUser <- getUserDetail
	}(req, l)

	wg.Wait()

	select {
	case val := <-chInterviewDetail:
		getInterviewDetail = val
	case val := <-chUser:
		getUserDetail = val
	case val := <-chInterviewDetailErr:
		interviewDetailErr = val
	case val := <-chUserErr:
		userErr = val
	}

	l.Info().Msg("go routine finished successfully")
	result := domains.GetInterviewDetailAll{
		GetInterviewDetail: getInterviewDetail,
		GetUserDetail:      getUserDetail,
		InterviewDetailErr: interviewDetailErr,
		UserErr:            userErr,
	}

	return result
}
