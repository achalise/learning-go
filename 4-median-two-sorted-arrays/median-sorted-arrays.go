package median_two_sorted_arrays

import (
	//"fmt"
	"math"
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	even := total % 2 == 0
	if even {
		return (findKth(total/2 + 1, nums1, nums2, 0, 0) + findKth(total/2, nums1, nums2, 0, 0)) * 0.5
	} else {
		return findKth(total/2 + 1, nums1, nums2, 0, 0)
	}
}
func findKth(k int, a1 []int, a2 []int, s1 int, s2 int) float64 {
	if s1  >= len(a1) {
		return float64(a2[s2 + k - 1])
	} else if s2  >= len(a2) {
		return float64(a1[s1 + k - 1])
	}
	if(k == 1) {
		return float64(math.Min(float64(a1[s1]), float64(a2[s2])))
	}
	m1 := s1 + k/2 -1
	fmt.Printf("The k %d m %d" , k, m1)
	m2 := s2 + k/2 -1
	var mid1 int
	var mid2 int
	if m1 > len(a1) {
		mid1 = 99999
	} else {
		fmt.Println("The m: ", m1)
		mid1 = a1[m1]
	}
	if m2 > len(a2) {
		mid2 = 99999
	} else {
		mid2 = a2[m2]
	}

	if(mid1 == mid2) {
		return float64(mid1)
	}
	if mid1 < mid2 {
		return findKth(k - k/2, a1, a2, m1 + 1, s2)
	} else {
		return findKth(k - k/2, a1, a2,s1, m2 + 1)
	}
}