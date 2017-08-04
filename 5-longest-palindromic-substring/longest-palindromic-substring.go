package longest_palindromic_substring


func longestPalindrome(s string) string {
	start := 0
	maxlength := 0
	l := len(s)
	if l == 1 {
		return s
	}
	for i := 0; i < l; i++ {
		lo1,length1 := extend(s, i, i);
		lo2,length2 := extend(s, i, i + 1);
		if (length1 > maxlength) {
			start = lo1
			maxlength = length1
		}
		if (length2 > maxlength) {
			start = lo2
			maxlength = length2
		}
	}
	return s[start: start + maxlength]
}

func extend(s string, i int, j int) (int, int) {
	start := 0
	length := 0
	for ;i >= 0 && j < len(s) && s[i] == s[j]; {
		start = i
		length = j - i + 1
		i--
		j++
	}
	return start, length
}