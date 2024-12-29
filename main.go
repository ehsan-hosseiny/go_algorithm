package main

import (
	"fmt"
	mergetwolists "github.com/ehsan-hosseiny/algorythm/mergetwolists"
)

func createLinkedList(nums []int) *mergetwolists.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &mergetwolists.ListNode{Val: nums[0]}
	current := head

	for _, val := range nums[1:] {
		current.Next = &mergetwolists.ListNode{Val: val}
		current = current.Next
	}

	return head
}

func printLinkedList(head *mergetwolists.ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() { // Create two linked lists
	list1 := createLinkedList([]int{1, 2, 4})
	list2 := createLinkedList([]int{1, 3, 4})

	// Merge the two linked lists
	mergedList := mergetwolists.mergeTwoLists(list1, list2)

	// Print the merged linked list
	printLinkedList(mergedList) // Output: 1 1 2 3 4 4

}
