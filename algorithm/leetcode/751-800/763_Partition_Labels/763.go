package _63_Partition_Labels

func partitionLabels(s string) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	lastIndex := make([]int, 26, 26)
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, lastIndex[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
