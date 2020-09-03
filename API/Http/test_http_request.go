package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	//"strings"
)

func main() {
	// GET请求
	// req_url :="http://httpbin.org/get"
	req_url := "http://www.baidu.com"
	client := &http.Client{}                         // 生成client客户端
	req, err := http.NewRequest("Get", req_url, nil) // 生成Request对象
	if err != nil {
		fmt.Println(err)
	}
	// 添加Header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close() // 设定关闭响应体
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// POST请求
	/*
		v := url.Values{}
		v.Set("username", "xxxx")
		v.Set("password", "xxxx")
		body := ioutil.NopCloser(strings.NewReader(v.Encode()))
		req, err := http.NewRequest("POST", "http://xxx.com/logindo", body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		fmt.Println(string(content))
	*/
}
