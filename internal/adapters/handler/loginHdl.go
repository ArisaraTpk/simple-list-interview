package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"simple-list-interview/internal/core/domains"
)

type LoginHdl interface {
	Login(c *gin.Context)
}

type loginHdl struct {
	svc domains.LoginSvc
}

func NewLoginHdl(svc domains.LoginSvc) LoginHdl {
	return &loginHdl{
		svc: svc,
	}
}

func (h loginHdl) Login(c *gin.Context) {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	var req domains.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
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
