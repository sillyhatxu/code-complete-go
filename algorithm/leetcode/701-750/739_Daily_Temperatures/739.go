package _39_Daily_Temperatures

func dailyTemperatures(temperatures []int) []int {
	var queue []int
	res := make([]int, len(temperatures), len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(queue) != 0 && temperatures[queue[len(queue)-1]] < temperatures[i] {
			idx := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			res[idx] = i - idx
		}
		queue = append(queue, i)
	}
	return res
}
