package ransom_note_383

func canConstruct(ransomNote string, magazine string) bool {
	// get freq of each character in magazine
	// iterate each character in ransome note
	// decrement freq dictionary if character can be found in magazine
	freq := make(map[rune]int)

	for _, key := range magazine {
		freq[key]++
	}

	for _, key := range ransomNote {
		val, keyExists := freq[key]
		if keyExists && val > 0 {
			freq[key]--
		} else {
			return false
		}
	}

	return true
}
