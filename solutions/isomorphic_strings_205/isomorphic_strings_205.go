package isomorphic_strings_205

func isIsomorphic(s string, t string) bool {
	forwardMap := make(map[rune]rune)
	backwardMap := make(map[rune]rune)

	for index, key := range s {
		forwardVal, forwardkeyExists := forwardMap[key]
		backwardVal, backwardkeyExists := backwardMap[rune(t[index])]

		if !forwardkeyExists && !backwardkeyExists {
			forwardMap[key] = rune(t[index])
			backwardMap[rune(t[index])] = key
		} else {
			if forwardkeyExists && forwardVal != rune(t[index]) {
				return false
			}

			if backwardkeyExists && backwardVal != key {
				return false
			}
		}
	}

	return true
}
