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

type UpdateInterviewDetailHdl interface {
	UpdateInterviewDetail(c *gin.Context)
}

type updateInterviewDetailHdl struct {
	svc domains.UpdateInterviewDetailSvc
}

func NewUpdateInterviewDetailHdl(svc domains.UpdateInterviewDetailSvc) UpdateInterviewDetailHdl {
	return &updateInterviewDetailHdl{
		svc: svc,
	}
}

func (h updateInterviewDetailHdl) UpdateInterviewDetail(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	interviewIdStr := c.Param("interviewId")
	interviewId, errParse := strconv.Atoi(interviewIdStr)
	if interviewIdStr == "" || errParse != nil {
		c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
			"errors": errors.ErrBadRequest,
		})
		return
	}

	var req domains.UpdateInterviewDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	req.InterviewId = int32(interviewId)
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
