package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	URL = "https://tieba.baidu.com/p/6051076813?red_tag=1573533731"
)

func main() {
	getEmail()
}

func getEmail() {
	resp, error := http.Get(URL)
	HandleError(error, "get url error")

	body := resp.Body

	bytes, err := ioutil.ReadAll(body)
	HandleError(err, "read body")

	html := string(bytes)

	re := regexp.MustCompile("(\\d+)@qq.com")

	arr := re.FindAllStringSubmatch(html, -1)

	if arr == nil {
		fmt.Println("not match")
		return
	}

	for _, val := range arr {
		fmt.Println(val)
	}

}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
