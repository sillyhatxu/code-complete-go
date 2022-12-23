package main

import "strings"

func strStrSample(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

func strStrOptimize(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	switch {
	case n == 0:
		return 0
	case n == 1:
		for i := 0; i < h; i++ {
			if string(haystack[i]) == needle {
				return i
			}
		}
		return -1
	case n == h:
		if haystack == needle {
			return 0
		}
		return -1
	case n > h:
		return -1
	default:
		for i := 0; i < h-n+1; i++ {
			if haystack[i:i+n] == needle {
				return i
			}
		}
		return -1
	}
}
