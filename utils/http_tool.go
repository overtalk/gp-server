package utils

import (
	"io/ioutil"
	"net/http"
)

var (
	tokenArgName = "token"
)

// GetToken : get token from gin.Context
func GetToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(tokenArgName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// GetRequestBody : get request body
func GetRequestBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}
