package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/protocol/v1"
	"github.com/golang/protobuf/proto"
)

var (
	postAddr = "http://127.0.0.1:5353/post"
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

	fmt.Printf("%d - %s", res.StatusCode, result)
}
