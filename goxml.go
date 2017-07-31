package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

type Vide struct {
	Durls         []Durl `xml:"durl"`
	Result        string `xml:"result"`
	Timelength    string `xml:"timelength"`
	Format        string `xml:"format"`
	AcceptFormat  string `xml:"accept_format"`
	AcceptQuality string `xml:"accept_quality"`
	From          string `xml:"from"`
	SeekParam     string `xml:"seek_param"`
	SeekType      string `xml:"seek_type"`
}

type Durl struct {
	Order  int64  `xml:"order"`
	Length string `xml:"length"`
	Url    string `xml:"url"`
}

type Result struct {
	Person []Person
}
type Person struct {
	Name      string
	Age       int
	Career    string
	Interests Interests
}
type Interests struct {
	Interest []string
}

func test() {
	input := `<?xml version="1.0" encoding="UTF-8"?>
<Persons>
    <Person>
        <Name>polaris</Name>
        <Age>28</Age>
        <Career>无业游民</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
    <Person>
        <Name>studygolang</Name>
        <Age>27</Age>
        <Career>码农</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
</Persons>`

	//inputReader := strings.NewReader(input)
	var result Result
	err := xml.Unmarshal([]byte(input), &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}

var (
	input = `<?xml version="1.0" encoding="UTF-8"?>
<video>
    <result a="test">suee</result>
    <timelength>2610988</timelength>
    <format>mp4</format>
    <accept_format><![CDATA[flv,hdmp4,mp4]]></accept_format>
    <accept_quality><![CDATA[3,2,1]]></accept_quality>
    <from><![CDATA[local]]></from>
    <seek_param><![CDATA[start]]></seek_param>
    <seek_type><![CDATA[second]]></seek_type>
    <durl>
        <order>1</order>
        <length>2610988</length>
        <size>93499980</size>
        <url><![CDATA[http://ws.acgvideo.com/3/f0/14578720-1.mp4?wsTime=1499857674&platform=pc&wsSecret2=caa6afab28f3bc4a69b19fe21edb72cd&oi=2014973402&rate=6]]></url>
    </durl>
</video>`
)

func main() {
	var t xml.Token
	var err error

	inputReader := strings.NewReader(input)
	decoder := xml.NewDecoder(inputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}
	}

	v := Vide{}
	err = xml.Unmarshal([]byte(input), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	//test()
}
