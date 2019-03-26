package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
)

var (
	serverAddr = "http://127.0.0.1:8080"
)

func main() {
	token := login()
	logout(token)
}

func login() string {
	var (
		postAddr = serverAddr + "/login"
	)

	req := &protocol.LoginReq{
		Account:  "tom0",
		Password: "tom0",
	}
	resp := &protocol.LoginResp{}

	reqByte, err := proto.Marshal(req)
	if err != nil {
		logger.Sugar.Fatalf("failed to marshal test request : %v", err)
	}

	client := &http.Client{}

	request, err := http.NewRequest("POST", postAddr, bytes.NewReader(reqByte))
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	res, err := client.Do(request)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logger.Sugar.Fatal(err)
		return ""
	}

	if err := proto.Unmarshal(result, resp); err != nil {
		logger.Sugar.Error(err)
		return ""
	}

	fmt.Printf("%d - %v", res.StatusCode, resp)
	return resp.Token
}

func logout(token string) {
	var (
		postAddr = serverAddr + "/logout"
	)

	client := &http.Client{}

	request, err := http.NewRequest("GET", postAddr, nil)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	// set cookie
	request.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})

	res, err := client.Do(request)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logger.Sugar.Fatal(err)
		return
	}

	resp := &protocol.LogoutResp{}
	if err := proto.Unmarshal(result, resp); err != nil {
		logger.Sugar.Error(err)
		return
	}

	fmt.Printf("%d - %d", res.StatusCode, resp.Status.Code)
}
