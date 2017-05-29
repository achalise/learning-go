package longest_substring

import "fmt"


//Longest substring with non-repeating characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	runes := []rune(s)

	length := 0
	maxlength := 0
	start := 0

	for i,rn := range runes {
		if val,ok := m[rn]; !ok || val < start {
			fmt.Println("length - , rn, i %d %c %d", length, rn, i)
			length += 1
			fmt.Println("length --, rn, i %d %c %d", length, rn, i)
		} else {
			fmt.Println("val, rn, i , length %d %c %d %d", val, rn, i, length)
			start = i
			length = i - val
			delete(m, rn)
		}
		if(length > maxlength) {
			maxlength = length;
		}
		m[rn] = i
	}
	return maxlength
}