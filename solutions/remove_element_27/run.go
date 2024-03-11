package remove_element_27

import "fmt"

func Run() {
	nums, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	res := removeElement(nums, val)
	fmt.Println(res)
}
