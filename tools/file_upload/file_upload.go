package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

var fileName = "1.zip"

func main() {
	// 上传文件POST
	// 下面构造一个文件buf作为POST的BODY
	filePath := "/Users/qinhan/go/src/github.com/qinhan-shu/gp-server/" + fileName
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, _ := w.CreateFormFile("uploadFile", fileName) //这里的uploadFile必须和服务器端的FormFile-name一致
	fd, _ := os.Open(filePath)
	defer fd.Close()
	io.Copy(fw, fd)
	w.Close()
	resp, _ := http.Post("http://0.0.0.0:8888/upload", w.FormDataContentType(), buf)
	helpRead(resp)
}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}
