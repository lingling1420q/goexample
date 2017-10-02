package main

import "fmt"

type Node struct {
	LeftChild   *Node
	RightChild  *Node
	ParentChild *Node
	Data        int
}

type SubTree struct {
	root  *Node
	nodes []*Node
}

func (self *SubTree) insert(node *Node) (err error) {
	self.nodes = append(self.nodes, node)
	n := self.root
	if n == nil {
		self.root = node
		return
	}
	for {
		if node.Data > n.Data {
			if n.RightChild == nil {
				n.RightChild = node
				node.ParentChild = n
				break
			} else {
				n = n.RightChild
			}
		} else if node.Data < n.Data {
			if n.LeftChild == nil {
				n.LeftChild = node
				node.ParentChild = n
				break
			} else {
				n = n.LeftChild
			}
		} else {
			fmt.Println("repeat value", node.Data)
			break
		}
	}
	return
}

func (self *SubTree) Len() int {
	return len(self.nodes)
}

func (self *SubTree) PreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.Data)
		self.PreOrder(node.LeftChild)
		self.PreOrder(node.RightChild)
	}
}

func (self *SubTree) InOrder(node *Node) {
	if node != nil {
		self.InOrder(node.LeftChild)
		fmt.Println(node.Data)
		self.InOrder(node.RightChild)
	}
}

func (self *SubTree) LastOrder(node *Node) {
	if node != nil {
		self.LastOrder(node.LeftChild)
		self.LastOrder(node.RightChild)
		fmt.Println(node.Data)
	}
}

func main() {
	subTree := &SubTree{}
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 33}
	node4 := &Node{Data: 4}
	node5 := &Node{Data: 15}
	node6 := &Node{Data: 0}

	subTree.insert(node1)
	subTree.insert(node2)
	subTree.insert(node3)
	subTree.insert(node4)
	subTree.insert(node5)
	subTree.insert(node6)

	fmt.Println(subTree.root)
	fmt.Println(subTree.Len())
	fmt.Println("---------PreOrder-----------")
	subTree.PreOrder(subTree.root)
	fmt.Println("---------InOrder-----------")
	subTree.InOrder(subTree.root)
	fmt.Println("---------LastOrder-----------")
	subTree.LastOrder(subTree.root)

	//     1
	// 0        2
	//     4        33
	//         15

}
