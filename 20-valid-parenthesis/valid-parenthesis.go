package valid_parenthesis


func isValid(s string) bool  {
    stack := make([]string, len(s) / 2 + 1)

	right := make(map[string]string);
	right["("] = ")";
	right["{"] = "}";
	right["["] = "]";


    length := 0
	for _, r := range s {
		if (length == 0 || right[stack[length - 1]] != string(r)) {
			if(length > len(s)/2) {
				return false;
			}
			stack[length] = string(r)
			length++
		} else {
			length = length - 1
		}
	}
	return length == 0;
}
