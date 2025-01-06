package main

import (
	"fmt"

	lengthOfLastWorld "github.com/ehsan-hosseiny/algorythm/length_of_last_world"
)

func main() {
	fmt.Println(lengthOfLastWorld.LengthOfLastWord("Hello World"))                 // Output: 5
	fmt.Println(lengthOfLastWorld.LengthOfLastWord("   fly me   to   the moon  ")) // Output: 4
	fmt.Println(lengthOfLastWorld.LengthOfLastWord("luffy is still joyboy"))       // Output: 6

}
