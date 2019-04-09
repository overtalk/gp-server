package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

var (
	tokenArgName = "token"
)

// GetToken : get token from gin.Context
func GetToken(r *http.Request) (string, error) {
	// FIXME: remote next line
	return tokenArgName, nil
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == tokenArgName {
			return cookie.Value, nil
		}
	}
	return "", errors.New("no token")
}

// GetRequestBody : get request body
func GetRequestBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

// MockHTTPReq : generate a http request for test
func MockHTTPReq(method, token string, req proto.Message) (*http.Request, error) {
	reqByte, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(method, "", bytes.NewReader(reqByte))
	if err != nil {
		return nil, err
	}

	r.AddCookie(&http.Cookie{
		Name:  tokenArgName,
		Value: token,
	})

	return r, nil
}

// GetReqAndToken : get token and protobuf data
func GetReqAndToken(c *http.Request) ([]byte, string, error) {
	// get data
	data, err := GetRequestBody(c)
	if err != nil {
		return nil, "", fmt.Errorf("missing request data")
	}
	// get token
	token, err := GetToken(c)
	if err != nil {
		return nil, "", fmt.Errorf("missing token")
	}
	return data, token, nil
}
