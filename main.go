package main

import (
	"fmt"

	searchinsertposition "github.com/ehsan-hosseiny/algorythm/search_insert_position"
)

func main() { // Create two linked lists

	nums := []int{1, 3, 5, 6}
	target := 7

	index := searchinsertposition.SearchInsert(nums, target) // Pass the correct arguments
	fmt.Println(index)                                       // Outputs: 2

}
