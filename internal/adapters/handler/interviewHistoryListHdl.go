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

type InterviewHistoryListHdl interface {
	InterviewHistoryList(c *gin.Context)
}

type interviewHistoryListHdl struct {
	svc domains.InterviewHistoryListSvc
}

func NewInterviewHistoryListHdl(svc domains.InterviewHistoryListSvc) InterviewHistoryListHdl {
	return &interviewHistoryListHdl{
		svc: svc,
	}
}

func (h interviewHistoryListHdl) InterviewHistoryList(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	interviewIdStr := c.Param("interviewId")
	interviewId, errParse := strconv.Atoi(interviewIdStr)
	if interviewIdStr == "" || errParse != nil {
		c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
			"errors": errors.ErrBadRequest,
		})
		return
	}

	req := domains.InterviewHistoryListReq{
		InterviewId: int32(interviewId),
	}

	res, err := h.svc.Execute(req, c, l)
	if err != nil {
		c.JSON(err.GetCode(), gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
