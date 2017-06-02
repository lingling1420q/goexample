package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"reflect"
)

//主要是这几个结构体
// type Selection struct {
//     Nodes    []*html.Node
//     document *Document
//     prevSel  *Selection
// }

// type Document struct {
//     *Selection
//     Url      *url.URL
//     rootNode *html.Node
// }

// type Node struct {
//     Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

//     Type      NodeType
//     DataAtom  atom.Atom
//     Data      string
//     Namespace string
//     Attr      []Attribute
// }

func Parse(i int, contentSelection *goquery.Selection) {
	info := contentSelection.Find(".title a")
	log.Println(info.Nodes[0])
	url, _ := info.Attr("href")
	log.Println("第", i+1, "个帖子的标题：", info.Text(), "url:", url)
}

func main() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".topics .topic").Each(Parse)
	//Find 方法返回Selection
	log.Println(".........................................................")
	//log.Println(doc.Find(".topic").Nodes[0])
	result, _ := doc.Find(".topic").Attr("class")
	log.Println(result)

	firstT := doc.Find(".topics .topic")
	html, _ := goquery.OuterHtml(firstT)
	t := reflect.TypeOf(firstT)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		//fmt.Println(m.Call())
		log.Printf("%6s: %v\n", m.Name, m.Type) //获取方法的名称和类型
	}

	log.Println(goquery.NodeName(firstT))
	log.Println(html)
	html, _ = firstT.Html()
	log.Println(html)
}
