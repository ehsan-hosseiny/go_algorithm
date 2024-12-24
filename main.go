package main

import (
	"fmt"

	LongestCommonPrefix "github.com/ehsan-hosseiny/algorythm/longest_common_prefix"
)

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(LongestCommonPrefix.LongestCommonPrefix(strs1))

}
