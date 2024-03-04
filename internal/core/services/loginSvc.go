package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"simple-list-interview/internal/core/domains"
	"simple-list-interview/internal/core/ports"
	"simple-list-interview/middleware/errors"
	"simple-list-interview/utils"
)

type loginSvc struct {
	validator *validator.Validate
	user      ports.UserRepo
}

func NewLoginSvc(validator *validator.Validate, user ports.UserRepo) domains.LoginSvc {
	return &loginSvc{
		validator: validator,
		user:      user,
	}
}

func (s loginSvc) Execute(req domains.LoginReq, ctx *gin.Context, l zerolog.Logger) (*domains.LoginRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	profile, err := s.user.FindUser(req.Username)
	if err != nil {
		l.Error().Err(err).Msg(fmt.Sprintf("FindUser errors %v", err))
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.ErrTechnical
	}

	isPasswordPass := s.validatePassword(req, profile)
	if !isPasswordPass {
		l.Error().Err(errors.ErrAuth).Msg("Password not match")
		return nil, errors.ErrAuth
	}

	jwtModify, errJwt := s.buildJWT(profile, l)
	if errJwt != nil {
		return nil, errJwt
	}

	return s.buildLoginRes(jwtModify), nil
}

func (s loginSvc) validate(req domains.LoginReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s loginSvc) buildLoginRes(jwt string) *domains.LoginRes {
	return &domains.LoginRes{
		Message: "Success",
		Token:   jwt,
	}
}

func (s loginSvc) validatePassword(req domains.LoginReq, user *ports.UserEntity) bool {
	if req.Password == user.Password {
		return true
	}
	return false
}

func (s loginSvc) buildJWT(user *ports.UserEntity, l zerolog.Logger) (string, *errors.APIError) {
	jwtModify, err := utils.GenerateJWTToken(user.UserId)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("GenerateJWTToken errors %v", err))
		return "", errors.ErrTechnical
	}

	return jwtModify, nil
}
