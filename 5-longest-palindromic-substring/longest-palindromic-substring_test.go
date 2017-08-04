package longest_palindromic_substring

import (
	"testing"
	"fmt"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	val := longestPalindrome("abcb")
	fmt.Printf(" The substring: %s", val)
}
