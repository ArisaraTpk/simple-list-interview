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

type GetInterviewDetailHdl interface {
	GetInterviewDetail(c *gin.Context)
}

type getInterviewDetailHdl struct {
	svc domains.GetInterviewDetailSvc
}

func NewGetInterviewDetailHdl(svc domains.GetInterviewDetailSvc) GetInterviewDetailHdl {
	return &getInterviewDetailHdl{
		svc: svc,
	}
}

func (h getInterviewDetailHdl) GetInterviewDetail(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	interviewIdStr := c.Param("interviewId")
	interviewId, errParse := strconv.Atoi(interviewIdStr)
	if interviewIdStr == "" || errParse != nil {
		c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
			"errors": errors.ErrBadRequest,
		})
		return
	}

	req := domains.GetInterviewDetailReq{
		InterviewId: int32(interviewId),
		UserId:      c.GetString("userId"),
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
