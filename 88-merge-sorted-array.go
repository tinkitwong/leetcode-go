package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	if m == 0 {
		copy(nums1, nums2)
		return nums1
	}

	i1 := m - 1
	i2 := n - 1

	for right := m + n - 1; right >= 0 && i2 >= 0; right-- {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[right] = nums1[i1]
			i1 -= 1
		} else {
			nums1[right] = nums2[i2]
			i2 -= 1
		}
	}
	return nums1
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1

	res := merge(nums1, m, nums2, n)
	fmt.Println(res)
}
