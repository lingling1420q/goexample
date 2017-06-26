package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
)

func main() {
	ch := rune(97)
	n := int('a')
	fmt.Printf("char: %c\n", ch)
	fmt.Printf("code: %d\n", n)

	var a int

	var b int32

	log.Println(a + int(b))

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))
	// output: e2c569be17396eca2a2e3c11578123ed

	// 直接使用md5 ew对象的Write方式也是一样的
	h2 := md5.New()
	h2.Write([]byte("The fog is getting thicker!"))
	h2.Write([]byte("And Leon's getting laaarger!"))
	fmt.Printf("%x\n", h2.Sum(nil))

	response, _ := http.Get("http://cache.video.qiyi.com/vps?tvid=686703800&vid=bd5787c5f44a63c7b62a7dad08b502ee&v=0&qypid=686703800_12&src=01012001010000000000&t=1497602104000&k_tag=1&k_uid=wogskflpeg6yievuh0kqdmwrwmz4vd1l&rs=1&vf=39c2002847bc9e9f5a06e7755baabf4a")

	bytes, _ := ioutil.ReadAll(response.Body)
	resp := string(bytes)
	log.Println(resp)
}
