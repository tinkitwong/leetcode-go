package summary_ranges_228

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)

	if n == 0 {
		return res
	}

	start := nums[0]
	for i := range n {
		if i == n-1 || nums[i]+1 != nums[i+1] {
			if nums[i] != start {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d", start))
			}
			if i != n-1 {
				start = nums[i+1]
			}
		}
	}

	return res
}
