package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Visited struct {
	Url       string
	Status    int
	VisitTime int `bson:"visitTime"`
}

func main() {
	src := "****"
	session, err := mgo.Dial(src) //连接数据库
	if err != nil {
		panic(err)
	}
	db := session.DB("crawl")     //数据库名称
	collection := db.C("visited") //如果该集合已经存在的话，则直接返回
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)
	//*****查询单条数据*******
	result := Visited{}
	err = collection.Find(bson.M{"url": "http://www.youtube.com/watch?v=o1KgYO1lrz4"}).One(&result)
	fmt.Println("Phone:", result.Url, result.VisitTime, result.Status)
}
