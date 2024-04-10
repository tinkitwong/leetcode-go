package plus_one_66

import "fmt"

// 123
// 124

// 129
// 130
func plusOne(digits []int) []int {
	carryover := 0
	digits[len(digits)-1]++

	// iterate from the back
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println(i, digits[i], carryover)
		sum := digits[i] + carryover
		digits[i] = sum % 10
		carryover = sum / 10
		fmt.Println("--", i, carryover, digits[i])
	}

	if carryover > 0 {
		digits = append([]int{carryover}, digits...)
	}

	return digits

}
