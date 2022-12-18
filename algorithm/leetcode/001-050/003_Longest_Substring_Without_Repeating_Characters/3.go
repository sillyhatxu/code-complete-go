package main

import (
	"fmt"
	"github.com/google/gxui/math"
)

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	check := make(map[byte]int)
	res, left := 0, 0
	for i := 0; i < len(s); i++ {
		if v, ok := check[s[i]]; ok && v >= left {
			left = check[s[i]] + 1
		}
		check[s[i]] = i
		res = max(res, i-left+1)
	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	maxLength := 0
	var currentArray []string
	for _, r := range s {
		c := string(r)
		for i, current := range currentArray {
			if current == c {
				currentArray = currentArray[i+1:]
			}
		}
		currentArray = append(currentArray, c)
		if maxLength < len(currentArray) {
			maxLength = len(currentArray)
		}
	}
	return maxLength
}

func Max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}

func lengthOfLongestSubstringOptimize(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[string]int)
	maxLength := 0
	j := 0
	for i, v := range s {
		c := string(v)
		if val, ok := m[c]; ok {
			j = Max(j, val+1)
		}
		m[c] = i
		maxLength = math.Max(maxLength, i-j+1)
	}
	return maxLength
}
func lengthOfLongestSubstringOptimizeBest(s string) int {
	var m [128]bool
	l := 0
	t := 0
	maxLength := 0
	for _, v := range s {
		c := v
		for m[c] {
			m[s[t]] = false
			t++
			l--
		}
		m[c] = true
		l++
		if maxLength < l {
			maxLength = l
		}
	}
	return maxLength
}

func main() {
	//test := []string{"abcabcbb", "bbbbb", "pwwkew", "ckilbkd"}
	test := []string{"pwwkew"}
	for _, src := range test {
		fmt.Println("lengthOfLongestSubstring : ", lengthOfLongestSubstring(src))
		fmt.Println("lengthOfLongestSubstringOptimize : ", lengthOfLongestSubstringOptimize(src))
		fmt.Println("lengthOfLongestSubstringOptimizeBest : ", lengthOfLongestSubstringOptimizeBest(src))
	}
}
