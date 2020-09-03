package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//dl :=[]byte("hello")
	resp, err := http.Get("http://www.baidu.com")
	check(err)
	defer resp.Body.Close()
	dl := []byte(resp.Status)
	// fmt.Println(dl)

	// write
	file_write_err := ioutil.WriteFile("test.txt", dl, 0644)
	check(file_write_err)

	// read
	data,file_read_err := ioutil.ReadFile("test.txt")
	check(file_read_err)
	fmt.Println(string(data))
}
