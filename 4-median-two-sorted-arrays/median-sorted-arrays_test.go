package median_two_sorted_arrays


import (
	"testing"
	"fmt"
)

func TestMedianSortedArrays(t *testing.T) {
	val := findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6})
	fmt.Printf(" The median: %f", val)
}
