package search_insert_position_35

import "fmt"

func Run() {
	nums := []int{1, 3, 5, 6}
	target := 7

	res := searchInsert(nums, target)
	fmt.Println(res)
}
