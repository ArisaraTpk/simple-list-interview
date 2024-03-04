package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"simple-list-interview/middleware/errors"
)

type LoginSvc interface {
	Execute(req LoginReq, ctx *gin.Context, l zerolog.Logger) (*LoginRes, *errors.APIError)
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
