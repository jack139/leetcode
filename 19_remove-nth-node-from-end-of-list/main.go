package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
 
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

 The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var buff [30]*ListNode

	d := 0
	for i:=head;i!=nil; {
		buff[d] = i
		d++
		i = i.Next
	}

	if d==n {
		return head.Next
	}

	buff[d-n-1].Next = buff[d-n].Next
	return head
}

func main(){
	var (
		n1 = ListNode{5, nil}
		n2 = ListNode{4, &n1}
		n3 = ListNode{3, &n2}
		n4 = ListNode{2, &n3}
		n5 = ListNode{1, &n4}
	)
	
	r := removeNthFromEnd(&n5, 2)
	i := r
	for i.Next!=nil {
		fmt.Println(i.Val)
		i = i.Next
	}
	fmt.Println(i.Val)
}
