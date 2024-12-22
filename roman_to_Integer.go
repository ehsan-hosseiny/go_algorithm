package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III")) // Output: 3
	fmt.Println(romanToInt("LVII")) // Output: 57
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994

}

func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	preValue := 0

	for i := len(s) - 1 ; i >= 0; i-- {
		currentValue := romanValues[s[i]]

		if currentValue < preValue {
            total -= currentValue
        } else {
            total += currentValue
        }
        preValue = currentValue
	}

	return total

}
