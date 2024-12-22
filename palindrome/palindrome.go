package palindrome

import "fmt"

func main() {
	// fmt.Println(isPalindrome(020))
	// fmt.Println(isPalindrome(121))  // true
	// fmt.Println(isPalindrome(-121)) // false
	// fmt.Println(isPalindrome(10))   // false
	fmt.Println(isPalindrome(1221)) // true
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x%10 == 0 && x != 0 {
		return false
	}

	reversed := 0
	original := x

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed

}
