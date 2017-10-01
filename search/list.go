package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	Left  *Node
	Right *Node
}

type List struct {
	Head  *Node
	nodes []*Node
}

func (list *List) insert(node *Node) (err error) {
	if list.Head == nil {
		list.Head = node
	} else {
		n := list.Head
		for {
			if n.Right == nil {
				n.Right = node
				node.Left = n
				break
			} else {
				n = n.Right
			}
		}
		list.nodes = append(list.nodes, node)
	}
	return
}

func (list *List) pop() (node *Node) {
	n := list.Head
	for {
		if n.Right == nil {
			list.nodes = list.nodes[:list.len()-1]
			return n
		} else {
			n = n.Right
		}
	}
}
func (list *List) len() int {
	return len(list.nodes)
}

func main() {
	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	list := &List{}
	list.insert(node1)
	list.insert(node3)
	list.insert(node4)
	list.insert(node2)
	n := list.Head
	for {
		fmt.Println(n.value)
		if n.Right == nil {
			break
		} else {
			n = n.Right
		}
	}
	fmt.Println("---------")
	fmt.Println(list.len())
	fmt.Println(list.pop().value)
	fmt.Println(list.len())
}
