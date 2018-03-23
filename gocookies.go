package main

import (
	//"github.com/axgle/mahonia"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

// data={
//     'redir': 'https://www.douban.com/',
//     'form_email':'18600189467',
//     'form_password':'linjie20061219'
//     #'login':u'登录'
// }

var (
	loginUrl string = "https://accounts.douban.com/login"
	indexUlr string = "https://www.douban.com/people/141268126/"
)

func addHeader(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
}

func RxOf(pattern string, content string, index int) (rcontent string) {
	re, _ := regexp.Compile(pattern)
	submatch := re.FindStringSubmatch(content)
	for i, v := range submatch {
		if i == index {
			rcontent = v
			break
		}
	}
	return
}

func R1(pattern string, content string) (rcontent string) {
	return RxOf(pattern, content, 1)
}

func httpDo() {
	client := &http.Client{}
	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar

	data := "redir=https://www.douban.com/&form_email=18600189467&form_password=linjie20061219"
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	addHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	h, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(h))
}

func main() {
	httpDo()
}
