package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], res) {
			res = res[:len(res)-1]
		}
	}
	return res
}

func longestCommonPrefixOriginal(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result, check, index, circulation := "", "", 0, true
	for circulation {
		for i, v := range strs {
			if i == 0 {
				if len(v)-1 < index {
					circulation = false
					break
				} else {
					check = string(v[index])
				}
			} else {
				if len(v)-1 < index {
					circulation = false
					break
				} else if check != string(v[index]) {
					circulation = false
					break
				}
			}
		}
		if circulation {
			result += check
			index++
		}
	}
	return result
}
