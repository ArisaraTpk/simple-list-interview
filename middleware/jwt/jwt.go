package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-list-interview/utils"
	"strings"
)

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

func CheckJWT(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	userId, err := utils.VerifyJWTToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: "bad jwt token",
		})
		return
	}

	c.Set("userId", userId)
	c.Next()
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
