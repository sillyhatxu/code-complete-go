package _25_Valid_Palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		} else if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		} else if unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right])) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
