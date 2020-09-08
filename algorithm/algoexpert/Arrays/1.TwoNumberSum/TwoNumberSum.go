package TwoNumberSum

func TwoNumberSum(array []int, target int) []int {
	check := make(map[int]int)
	for _, value := range array {
		index, ok := check[value]
		if ok {
			return []int{target - value, value}
		}
		check[target-value] = index
	}
	return []int{}
}
