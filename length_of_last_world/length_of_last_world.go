package lengthOfLastWorld

import "strings"

func LengthOfLastWord(s string) int {
	//Trim ant training spaces
	s = strings.TrimSpace(s)

	// Split the string by spaces
	words := strings.Split(s, " ")

	//Return the length of the last world
	return len(words[len(words)-1])

}
