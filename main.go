package main

import (
	"fmt"

	romantointeger "github.com/ehsan-hosseiny/algorythm/roman_to_integer"
)

func main() {
	
	fmt.Println(romantointeger.RomanToInt("III"))     // Output: 3
	fmt.Println(romantointeger.RomanToInt("LVII"))    // Output: 57
	fmt.Println(romantointeger.RomanToInt("MCMXCIV")) // Output: 1994

}
