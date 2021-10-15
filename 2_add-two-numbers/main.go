package main

import (
	"fmt"
)


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var znode *ListNode
	var x, y, h, m int

	r := l2 // 第一个节点是l2的第一个
	last := r
	for ;l1!=l2; { // 只有两个都是nil才相等
		if l1==nil { x=0 } else { x=l1.Val; znode=l1; l1=l1.Next }
		if l2==nil { y=0 } else { y=l2.Val; znode=l2; l2=l2.Next }
		h = x+y+m
		if h>9 {
			znode.Val = h - 10 // 当前位
			m = 1 // 进位
		} else {
			znode.Val = h
			m =0
		}
		last.Next = znode
		last = znode
	}
	last.Next = nil
	if m>0 {
		last.Next = &ListNode{m, nil}
	}
	return r
}

func printList(l *ListNode){
	for next:=l;next!=nil;next=next.Next {
		fmt.Printf("%d ", next.Val)
	}
	fmt.Println()
}

func main(){
	var l1 = []ListNode{{2,nil},{4,nil},{3,nil},{9,nil}}
	var l2 = []ListNode{{5,nil},{6,nil},{7,nil}}

	l1[0].Next = &l1[1]
	l1[1].Next = &l1[2]
	l1[2].Next = &l1[3]

	l2[0].Next = &l2[1]
	l2[1].Next = &l2[2]

	printList(&l1[0])
	printList(&l2[0])

	r := addTwoNumbers(&l1[0], &l2[0])
	printList(r)
}