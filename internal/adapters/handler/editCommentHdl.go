package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"simple-list-interview/internal/core/domains"
	"simple-list-interview/middleware/errors"
	"strconv"
)

type EditCommentHdl interface {
	EditComment(c *gin.Context)
}

type editCommentHdl struct {
	svc domains.EditCommentSvc
}

func NewEditCommentHdl(svc domains.EditCommentSvc) EditCommentHdl {
	return &editCommentHdl{
		svc: svc,
	}
}

func (h editCommentHdl) EditComment(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	commentIdStr := c.Param("commentId")
	commentId, errParse := strconv.Atoi(commentIdStr)
	if commentIdStr == "" || errParse != nil {
		c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
			"errors": errors.ErrBadRequest,
		})
		return
	}

	req := domains.EditCommentReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	req.CommentId = int32(commentId)
	req.UserId = c.GetString("userId")

	res, err := h.svc.Execute(req, c, l)
	if err != nil {
		c.JSON(err.GetCode(), gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
