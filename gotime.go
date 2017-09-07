package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func getNow() {
	// 获取当前时间，返回time.Time对象
	fmt.Println(time.Now())
	// output: 2016-07-27 08:57:46.53277327 +0800 CST
	// 其中CST可视为美国，澳大利亚，古巴或中国的标准时间
	// +0800表示比UTC时间快8个小时

	// 获取当前时间戳
	fmt.Println(time.Now().Unix())
	// 精确到纳秒，通过纳秒就可以计算出毫秒和微妙
	fmt.Println(time.Now().UnixNano())
	// output:
	//    1469581066
	//    1469581438172080471
}

func formatUnixTime() {
	// 获取当前时间，进行格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:57:46

	// 指定的时间进行格式化
	fmt.Println(time.Unix(1469579899, 0).Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:38:19
}

func getYear() {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(1469579899, 0)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// output: 2016-7-27 8:38:19
}

//时间字符串转换时间戳
// 将2016-07-27 08:46:15这样的时间字符串转换时间戳
func strToUnix() {
	// 先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
	// 后面就可以用time.Time对象的函数Unix进行获取
	t2, err := time.Parse("2006-01-02 15:04", "2017-02-26 03:42")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t2)
	fmt.Println(t2.Unix())
	// output:
	//     2016-07-27 08:46:15 +0000 UTC
	//     1469609175
}

func getDayStartUnix() {
	t := time.Unix(1469581066, 0).Format("2006-01-02")
	sts, err := time.Parse("2006-01-02", t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sts.Unix())
	// output: 1469577600
}

var (
	ErrParse = errors.New("Parse time error")
)

var (
	timeLayoutsForParse = []string{
		"20060102150405",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.RFC3339,
		time.RFC3339Nano,
		"20060102",
		"2006-01-02",                         // RFC 3339
		"2006-01-02 15:04",                   // RFC 3339 with minutes
		"2006-01-02 15:04:05",                // RFC 3339 with seconds
		"2006-01-02 15:04:05-07:00",          // RFC 3339 with seconds and timezone
		"2006-01-02T15Z0700",                 // ISO8601 with hour
		"2006-01-02T15:04Z0700",              // ISO8601 with minutes
		"2006-01-02T15:04:05Z0700",           // ISO8601 with seconds
		"2006-01-02T15:04:05.999999999Z0700", // ISO8601 with nanoseconds
	}
)

func TryParse(s string) (time.Time, error) {
	for _, layout := range timeLayoutsForParse {
		r, err := time.Parse(layout, s)
		if err == nil {
			return r, nil
		}
	}
	return time.Time{}, ErrParse
}

func MustParse(s string) time.Time {
	t, err := TryParse(s)
	if err != nil {
		panic(err)
	}
	return t
}

// returns a formatted string of `time.RFC1123` format.
func TimeStrToRFC1123(str string) string {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t, err = time.Parse(time.RFC1123, str)
		if err != nil {
			panic("Time format invalid. The time format must be time.RFC3339 or time.RFC1123")
		}
	}
	return t.Format(time.RFC1123)
}

// returns a utc string of a time instance.
func TimeToUTCStr(t time.Time) string {
	format := time.RFC3339 // 2006-01-02T15:04:05Z07:00
	return t.UTC().Format(format)
}

func TimestampToHourStr(ts int) string {
	//格式化为字符串,tm为Time类型
	tm := time.Unix(int64(ts), 0)
	return tm.Format("3小时4分5秒")
}

func main() {
	// getNow()
	// formatUnixTime()
	strToUnix()
	fmt.Println(TimestampToHourStr(4000))
	a := time.Now().Format(time.RFC1123)
	fmt.Println(a)

	b := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Println(b.Add(1 * time.Second).Before(time.Now()))
	fmt.Println(b.Add(1 * time.Second).After(time.Now()))
}
