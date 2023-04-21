package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	left := 0
	right := len(s) - 1

	for left < right {
		if !isAlphaNum(rune(s[left])) {
			left++
			continue
		}
		if !isAlphaNum(rune(s[right])) {
			right--
			continue
		}
		if lower(s[left]) != lower(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func lower(b byte) string {
	return strings.ToLower(string(b))
}

func isAlphaNum(u rune) bool {
	return unicode.IsLetter(u) || unicode.IsDigit(u)
}
