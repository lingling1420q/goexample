package main

import (
	"fmt"
)

type SLink struct {
	value int
	next  *SLink
}

func print(head *SLink) {
	tmp := head
	for tmp.next != nil {
		fmt.Print(tmp.value, " ")
		tmp = tmp.next
	}
	fmt.Print(tmp.value, "\n")
}

func reverse01(head *SLink) (tail *SLink) {
	if head == nil || head.next == nil {
		tail = head
	} else {
		tail = reverse01(head.next)
		head.next.next = head
		head.next = nil
	}
	return
}

func reverse02(head *SLink) (tail *SLink) {
	if head == nil || head.next == nil {
		tail = head
	} else {
		pre := head
		cur := head.next
		tmp := head.next.next
		for cur != nil {
			tmp = cur.next
			cur.next = pre
			pre = cur
			cur = tmp
		}
		head.next = nil
		tail = pre
	}
	return
}

func main() {
	node1 := &SLink{value: 1}
	node2 := &SLink{value: 2}
	node3 := &SLink{value: 3}
	node4 := &SLink{value: 4}
	node5 := &SLink{value: 5}
	node6 := &SLink{value: 6}
	node7 := &SLink{value: 7}
	node8 := &SLink{value: 8}
	node9 := &SLink{value: 9}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9
	print(node1)
	fmt.Println("原始版")
	n1 := reverse01(node1)
	print(n1)
	fmt.Println("递归版")
	n2 := reverse02(n1)
	print(n2)
	fmt.Println("for版")
}
