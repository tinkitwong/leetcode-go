package search_insert_position_35

func searchInsert(nums []int, target int) int {
	// binary search to find target
	left := 0
	right := len(nums) - 1

	// find target if exists
	for right >= left {
		mid := (left + right) / 2

		// move find algo to the left <-
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}
