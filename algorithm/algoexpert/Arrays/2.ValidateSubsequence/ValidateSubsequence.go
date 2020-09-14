package ValidateSubsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	check := make(map[int]int)
	for _, v := range array {
		_, ok := check[v]
		if ok {
			check[v] = check[v] + 1
		} else {
			check[v] = 1
		}
	}
	for _, v := range sequence {
		result, ok := check[v]
		if !ok || result == 0 {
			return false
		}
		check[v] = check[v] - 1
	}
	return true
}
