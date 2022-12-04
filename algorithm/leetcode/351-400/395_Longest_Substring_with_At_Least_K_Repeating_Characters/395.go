package _95_Longest_Substring_with_At_Least_K_Repeating_Characters

func longestSubstring(s string, k int) int {
	return helper([]byte(s), 0, len(s), k)
}

func helper(letters []byte, start, end, k int) int {
	if end-start < k {
		return 0
	}
	frequency := [26]int{}
	for i := start; i < end; i++ {
		frequency[letters[i]-'a']++
	}
	for i := byte(0); i < 26; i++ {
		//if s = aaabbbcaabb; k = 2; we will find c letter
		if frequency[i] > 0 && frequency[i] < k {
			for j := start; j <= end; j++ {
				if i == letters[j]-'a' { //letter c
					left := helper(letters, start, j, k)
					right := helper(letters, j+1, end, k)
					return max(left, right)
				}
			}
		}
	}
	return end - start
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
