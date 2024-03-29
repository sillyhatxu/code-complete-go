package _40_Longest_Substring_with_At_Most_K_Distinct_Characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	indexMap := make(map[byte]int)
	res, left := 0, 0
	for right := 0; right < len(s); right++ {
		indexMap[s[right]]++
		for len(indexMap) > k {
			indexMap[s[left]]--
			if indexMap[s[left]] == 0 {
				delete(indexMap, s[left])
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func lengthOfLongestSubstringKDistinct1(s string, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	indexMap := make(map[byte]int)
	start, size := 0, 0
	for i := 0; i < len(s); i++ {
		indexMap[s[i]] = i
		if len(indexMap) > k {
			delIndex := i
			for _, index := range indexMap {
				delIndex = min(delIndex, index)
			}
			start = delIndex + 1
			delete(indexMap, s[delIndex])
		}
		size = max(size, i-start+1)
	}
	return size
}
