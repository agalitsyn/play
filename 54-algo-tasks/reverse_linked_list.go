// Level: easy
//
// Source: https://leetcode.com/problems/reverse-linked-list/solution/
//
// Reverse a singly linked list.
//
// Example:
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
//
// Solution: https://www.geeksforgeeks.org/reverse-a-linked-list/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(uintptr(0)))
	res := Solve(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
	)
	fmt.Printf("%+v\n", res)
}

func iter(node *ListNode) *ListNode {
	if node.Next != nil {
		return node.Next
	}
	return nil
}

// 1. Initialize three pointers prev as NULL, curr as head and next as NULL.
// 2. Iterate through the linked list. In loop, do following.
// Before changing next of current,
// store next node
// next = curr->next
// Now change next of current
// This is where actual reversing happens
// curr->next = prev
// Move prev and curr one step forward
// prev = curr
// curr = next
func Solve(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
