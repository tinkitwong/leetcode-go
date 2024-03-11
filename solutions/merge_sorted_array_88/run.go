package merge_sorted_array_88

import "fmt"

func Run() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1

	res := merge(nums1, m, nums2, n)
	fmt.Println(res)
}
