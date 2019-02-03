// PostFile2RD project main.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var gPostMsgUrl string ="" //Post File Url Addr

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, err
}

func getArgs() []string {
	var inputList []string
	args := os.Args
	for i := 0; i < len(args); i++ {
		inputList = append(inputList, args[i])
		fmt.Println(args[i])
	}
	return inputList
}
func testRun() {
	path := "D:\\bg.gif"
	extraParams := map[string]string{
		"description":   "产品建议333微信可以批改作业",
		"teacherName":   "绿巨人",
		"teacherNumber": "ga770546248",
		"creationTime":  "2019.02.02 18:09:20",
		"version":       "test版本2.0.0.0",
		"type":          "1",
	}
	request, err := newfileUploadRequest(gPostMsgUrl, extraParams, "userfile", path)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)

		fmt.Println(body)
		if 200 != resp.StatusCode {
			fmt.Println("发送失败")
		} else {
			fmt.Println("发送成功")
		}
	}
}
func PublishRun() {
	inputList := getArgs()
	if len(inputList) != 8 {
		fmt.Printf("请输入描述、辅导老师账号、辅导老师Number、反馈创建时间、版本号、描述类型、附件路径!\n")
		return
	}
	path := inputList[7]

	extraParams := map[string]string{
		"description":   inputList[1],
		"teacherName":   inputList[2],
		"teacherNumber": inputList[3],
		"creationTime":  inputList[4],
		"version":       inputList[5],
		"type":          inputList[6],
	}
	request, err := newfileUploadRequest(gPostMsgUrl, extraParams, "userfile", path)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)

		fmt.Println(body)
		if 200 != resp.StatusCode {
			fmt.Println("发送失败")
		} else {
			fmt.Println("发送成功")
		}
	}
}
func main() {
	PublishRun()
}
