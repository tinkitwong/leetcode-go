package valid_palindrome

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// lowercase s
	// remove non alpha numeric characters
	// checking 2 pointers for palindrome
	reg, _ := regexp.Compile(`[^a-zA-Z0-9\\s]+`)
	regex_str := reg.ReplaceAllString(s, "")
	str := strings.ToLower(regex_str)

	fmt.Println(str)

	l := 0
	r := len(str) - 1

	for l <= r {
		if str[l] == str[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}

func isPalindromeII(s string) bool {
	// takes in unicode code point
	f := func(r rune) rune {
		// if not letter and not number
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		// return only alpha numerics in lower case
		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
