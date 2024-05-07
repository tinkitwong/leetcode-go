package remove_duplicates_from_sorted_array_26

func removeDuplicates(nums []int) int {
	l := 0
	for l < len(nums)-1 {
		// check if next element is the same as current
		// if same, then remove
		if nums[l+1] == nums[l] {
			nums = append(nums[:l+1], nums[l+2:]...)
		} else {
			l++
		}
	}

	return len(nums)
}
