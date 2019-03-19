package utils

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var (
	tokenArgName = "token"
)

// GetToken : get token from gin.Context
func GetToken(c *gin.Context) (string, error) {
	cookie, err := c.Request.Cookie(tokenArgName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// GetRequestBody : get request body
func GetRequestBody(c *gin.Context) ([]byte, error) {
	return ioutil.ReadAll(c.Request.Body)
}
