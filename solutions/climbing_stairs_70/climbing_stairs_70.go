package climbing_stairs_70

func climbStairs(n int) int {

	// set base case
	res := []int{1, 2}

	if n < 3 {
		return res[n-1]
	}

	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}

	return res[n-1]
}
