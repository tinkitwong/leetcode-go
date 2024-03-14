package is_subsequence_392

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := 0
	for i < len(t) {
		if j == len(s) {
			return true
		}
		// match the current element of subsequence
		// increment j pointer up
		if t[i] == s[j] {
			j++
		}

		i++
	}

	return j == len(s)
}
