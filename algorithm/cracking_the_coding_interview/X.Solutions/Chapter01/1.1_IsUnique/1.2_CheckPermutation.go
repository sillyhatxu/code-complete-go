package main

func CheckPermutation(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}
	check := make([]int32, 128)
	for _, c := range input1 {
		check[c]++
	}
	for _, c := range input2 {
		check[c]--
		if check[c] < 0 {
			return false
		}
	}
	return true
}
