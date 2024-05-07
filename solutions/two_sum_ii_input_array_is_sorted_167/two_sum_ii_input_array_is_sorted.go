package twosumiiinputarrayissorted167

func TwoSum(numbers []int, target int) []int {
	// because it is sorted, there is a way to tell
	// if the current two values are greater/lesser than target

	// if current sum > target: need to lower the sum
	// ie. we need to move right pointer to left

	// if current sum < target: need to increase the sum
	// ie. we need to move left pointer to right

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}
