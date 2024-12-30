package main

import (
	"fmt"

	findfirstoccurrence "github.com/ehsan-hosseiny/algorythm/find_index_first_occurrence"
)

func main() { // Create two linked lists

	haystack2 := "leetcodeleet"
	needle2 := "le"

	indices2 := findfirstoccurrence.StrStr(haystack2, needle2)
	fmt.Println(indices2) // Output: []

}
