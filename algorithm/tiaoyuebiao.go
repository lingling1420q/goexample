package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	value int
}

type SkipList struct {
	items *Item
	next  *SkipList
	//index [][]*Item
}

type LikeNode struct {
	item *Item
	next *LikeNode
}

func (self *LikeNode) add(item *Item) {
	var node LikeNode
	if self.item.value > item.value {

	}
}

func (self *LikeNode) del(item *Item) {

}

func main() {
	rand.Seed(time.Now().UnixNano())
	likeNode := LikeNode{num: 1}
	for i := 1; i < 100; i++ {
		it := rand.Intn(100)
		item := &Item{value: it}
		//fmt.Println(item)
		likeNode.add(item)
	}

	fmt.Println(likeNode)
}
