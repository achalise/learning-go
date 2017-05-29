package two_sum

func twoSum(input []int, target int) []int {
	indicesMap := make(map[int]int)
	for i:= 0; i < len(input); i++ {
		diff := target - input[i]
		if _, ok := indicesMap[diff]; ok {
			return []int{i, indicesMap[diff]}
		} else {
			indicesMap[input[i]] = i
		}
	}
  return nil;
}

func bruteForce(input []int, target int)[]int {
    for i := 0; i < len(input); i++ {
			 for j := 0; j < len(input); j++ {
				 if( i == j) {
					 continue
				 }
				 if(input[i] + input[j] == target) {
					 return ([]int{input[i], input[j]})
				 }
			 }
    }
    return []int{0,0}
}


