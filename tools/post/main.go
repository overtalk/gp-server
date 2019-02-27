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
	postAddr = "http://127.0.0.1:5353/example"
)

func main() {
	req := &protocol.TestRequest{
		A: "aaa",
		B: "bbb",
	}

	reqByte, err := proto.Marshal(req)
	if err != nil {
		logger.Sugar.Fatalf("failed to marshal test request : %v", err)
	}

	body := bytes.NewBuffer(reqByte)
	res, err := http.Post(postAddr, "application/binary;charset=utf-8", body)
	if err != nil {
		logger.Sugar.Fatal(err)
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logger.Sugar.Fatal(err)
		return
	}

	resp := &protocol.TestResponse{}
	if err := proto.Unmarshal(result, resp); err != nil {
		logger.Sugar.Error(err)
		return
	}

	fmt.Printf("%d - %v", res.StatusCode, resp)
}
