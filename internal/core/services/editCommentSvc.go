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

type editCommentSvc struct {
	validator *validator.Validate
	comments  ports.CommentRepo
}

func NewEditCommentSvc(validator *validator.Validate, comments ports.CommentRepo) domains.EditCommentSvc {
	return &editCommentSvc{
		validator: validator,
		comments:  comments,
	}
}

func (s editCommentSvc) Execute(req domains.EditCommentReq, ctx *gin.Context, l zerolog.Logger) (*domains.EditCommentRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	commentDetail, err := s.comments.FindCommentDetail(ports.FindCommentDetailReq{CommentId: req.CommentId})
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("FindCommentDetail errors %v", err))
		return nil, errors.ErrNotFound
	}

	if req.UserId != commentDetail.CreatedBy {
		l.Error().
			Msg("User don't have permission to edit this comment")
		return nil, errors.ErrBadRequest
	}

	updateComment := ports.UpdateCommentDescriptionReq{
		CommentId:   req.CommentId,
		Description: req.Description,
	}
	err = s.comments.UpdateCommentDescription(updateComment)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("UpdateCommentDescription errors %v", err))
		return nil, errors.ErrTechnical
	}

	return &domains.EditCommentRes{Message: "Success"}, nil
}

func (s editCommentSvc) validate(req domains.EditCommentReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}
