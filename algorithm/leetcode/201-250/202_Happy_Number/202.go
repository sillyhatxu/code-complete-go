package _02_Happy_Number

func isHappy(n int) bool {
	history := make(map[int]struct{})
	for n != 1 {
		temp := 0
		for n != 0 {
			num := n % 10
			temp += num * num
			n = n / 10
		}
		if _, ok := history[temp]; ok {
			return false
		}
		history[temp] = struct{}{}
	}
	return true
}
