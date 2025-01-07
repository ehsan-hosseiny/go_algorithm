package main

import (
	"fmt"

	plusOne "github.com/ehsan-hosseiny/algorythm/plus_one"
)

func main() {

	// Example usage:
	digits1 := []int{1, 2, 3}
	digits2 := []int{4, 3, 2, 1}
	digits3 := []int{9}

	fmt.Println(plusOne.PlusOne(digits1)) // Output: [1, 2, 4]
	fmt.Println(plusOne.PlusOne(digits2)) // Output: [4, 3, 2, 2]
	fmt.Println(plusOne.PlusOne(digits3)) // Output: [1, 0]

}
