package _25_Valid_Palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		} else if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		} else if unicode.ToLower(rune(s[l])) == unicode.ToLower(rune(s[r])) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome1(s string) bool {
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
