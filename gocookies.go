package main

import (
	//"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	//"time"
	//"strings"
)

func httpDo() {
	client := &http.Client{}
	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar

	req, err := http.NewRequest("GET", "http://www.66ip.cn/areaindex_1/index.html", nil)
	if err != nil {
		// handle error
	}

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "name=anny")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Host", "Host:www.66ip.cn")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Set("Referer", "http://www.66ip.cn/areaindex_1/index.html")
	req.Header.Set("Cookie", "__cfduid=dec031f0ed6ab5afc64978bfaa48a9e0a1496720691; UM_distinctid=15c7b8161ab1-06d3d1f923b304-3062750a-fa000-15c7b8161ac32e; cf_clearance=72e6dbcd3682413fd4692076abf4da6eff2e44c8-1496731093-3600; CNZZDATA1253901093=1678483328-1496717997-http%253A%252F%252Fwww.66ip.cn%252F%7C1496728955; Hm_lvt_1761fabf3c988e7f04bec51acd4073f4=1496720696; Hm_lpvt_1761fabf3c988e7f04bec51acd4073f4=1496732616")

	for _, obj := range req.Cookies() {
		log.Println(obj.Name, ":", obj.Value, obj.Domain)
	}

	log.Println(req.Header)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	html := string(body)
	//html = mahonia.NewDecoder("gbk").ConvertString(html)
	log.Println(len(html))
	log.Println(html)
}

func main() {
	httpDo()
}
