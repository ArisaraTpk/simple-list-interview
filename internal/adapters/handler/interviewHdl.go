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

type InterviewListHdl interface {
	InterviewList(c *gin.Context)
}

type interviewListHdl struct {
	svc domains.InterviewListSvc
}

func NewInterviewListHdl(svc domains.InterviewListSvc) InterviewListHdl {
	return &interviewListHdl{
		svc: svc,
	}
}

func (h interviewListHdl) InterviewList(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	var lastOrderId, size int
	var errValidate error
	lastOrderIdStr, isFound := c.GetQuery("lastItemOrder")
	if !isFound {
		lastOrderId = 0
	} else {
		lastOrderId, errValidate = strconv.Atoi(lastOrderIdStr)
		if errValidate != nil {
			c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
				"errors": errors.ErrBadRequest,
			})
			return
		}
	}

	sizeStr, isFound := c.GetQuery("size")
	if !isFound {
		size = 10
	} else {
		size, errValidate = strconv.Atoi(sizeStr)
		if errValidate != nil {
			c.JSON(errors.ErrBadRequest.GetCode(), gin.H{
				"errors": errors.ErrBadRequest,
			})
			return
		}
	}

	req := domains.InterviewListReq{
		LastItemOrder: lastOrderId,
		Size:          size,
		UserId:        c.GetString("userId"),
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
