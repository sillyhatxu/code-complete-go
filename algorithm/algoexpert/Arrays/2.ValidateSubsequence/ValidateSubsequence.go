package ValidateSubsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) < len(sequence) {
		return false
	}
	index := 0
	for _, v := range array {
		if v == sequence[index] {
			index++
		}
		if index == len(sequence) {
			break
		}
	}
	return index == len(sequence)
}
