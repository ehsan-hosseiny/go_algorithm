package LongestCommonPrefix

import "sort"

func LongestCommonPrefix(strs []string) string {
	// if len(strs) == 0 {
	// 	return ""
	// }

	// // ["flower","flow","flight"]

	// for i := 1; i <= len(strs[0]); i++ {
	// 	for _, str := range strs {
	// 		if strings.HasPrefix(str, strs[0][:i]) {
	// 			continue
	// 		} else {
	// 			if i == 1 {
	// 				return ""
	// 			}
	// 			return strs[0][:i-1]
	// 		}
	// 	}
	// 	// Deal with the situation when the first string is the common prefix
	// 	if i == len(strs[0]) {
	// 		return strs[0]
	// 	}
	// }
	// return ""

	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix

}
