package two_sum

import (
   "testing"
   "fmt"
)

func TestBruteForce(t *testing.T) {
	res := bruteForce([]int{1,2,3,4,8}, 6)
	fmt.Printf(" The values %d : %d ", res[0], res[1])
}

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{1,2,3,4,8}, 6);
	fmt.Printf(" The values %d : %d ", res[0], res[1])
}
