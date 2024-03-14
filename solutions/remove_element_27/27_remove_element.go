package remove_element_27

func removeElement(nums []int, val int) int {
	count := 0
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == val {
			// remove element
			// add filler element at the end
			// count ++

			nums = append(nums[:l], nums[l+1:]...)
			nums = append(nums, 0)
			r -= 1
			count += 1
		} else {
			l += 1
		}
	}

	return len(nums) - count
}

func removeElementII(nums []int, val int) int {
	// pointer for valid part of slice
	j := 0
	for i := 0; i < len(nums); i++ {
		// nums[i] != val
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return j
}
