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
	postAddr = "http://127.0.0.1:9999/login"
)

func main() {
	req := &protocol.LoginReq{
		Account:  "aaa",
		Password: "bbb",
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

	// set cookie
	// request.AddCookie(&http.Cookie{
	// 	Name:  "token",
	// 	Value: "aaa",
	// })

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

	if err := proto.Unmarshal(result, resp); err != nil {
		logger.Sugar.Error(err)
		return
	}

	fmt.Printf("%d - %v", res.StatusCode, resp)
}
