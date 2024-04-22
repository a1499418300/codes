package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	url := "http://example.com/upload"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open("file_path")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, err := writer.CreateFormFile("file", "file_path")
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(part1, file)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println("File uploaded successfully!")
}
