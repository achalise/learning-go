package longest_substring

import (
	"testing"
	"fmt"
)

func TestLengthOfLongestSubstring(t *testing.T) {
   len := lengthOfLongestSubstring("dvdf")
	fmt.Printf(" The length of longest substring with nonrepeating characters %d", len)
}
