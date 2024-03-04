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

type commentListSvc struct {
	validator *validator.Validate
	comments  ports.CommentRepo
}

func NewCommentListSvc(validator *validator.Validate, comments ports.CommentRepo) domains.CommentListSvc {
	return &commentListSvc{
		validator: validator,
		comments:  comments,
	}
}

func (s commentListSvc) Execute(req domains.CommentListReq, ctx *gin.Context, l zerolog.Logger) (*domains.CommentListRes, *errors.APIError) {
	if err := s.validate(req, l); err != nil {
		return nil, err
	}

	reqComments := ports.FindCommentListReq{
		InterviewId: req.InterviewId,
	}
	commentList, err := s.comments.FindCommentList(reqComments)
	if err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("call FindInterviewDetail failed %v", err))
		return nil, errors.ErrTechnical
	}
	return s.buildCommentListRes(commentList), nil
}

func (s commentListSvc) validate(req domains.CommentListReq, l zerolog.Logger) *errors.APIError {
	if err := s.validator.Struct(req); err != nil {
		l.Error().
			Err(err).
			Msg(fmt.Sprintf("validate errors %v", err))
		return errors.ErrBadRequest
	}
	return nil
}

func (s commentListSvc) buildCommentListRes(comments []ports.CommentEntity) *domains.CommentListRes {
	commentResList := []domains.Comment{}
	for _, comment := range comments {
		data := domains.Comment{
			CommentId:   comment.CommentId,
			Description: comment.Description,
			CreatedBy:   comment.CreatedBy,
			CreatedAt:   comment.CreatedAt,
			UpdatedAt:   comment.UpdatedAt,
		}
		commentResList = append(commentResList, data)
	}

	return &domains.CommentListRes{
		Comments: commentResList,
	}
}
