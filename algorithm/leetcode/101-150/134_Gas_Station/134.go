package _34_Gas_Station

func canCompleteCircuit(gas []int, cost []int) int {
	surplusGas := 0
	currentGas := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
		surplusGas += gas[i] - cost[i]
	}
	if surplusGas >= 0 {
		return start
	}
	return -1
}

//canCompleteCircuitBackups 暴力解法
func canCompleteCircuitBackups(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if helper(i, gas, cost) {
			return i
		}
	}
	return -1
}

func helper(start int, gas []int, cost []int) bool {
	currentGas := gas[start]
	index := start
	for {
		currentCost := cost[index]
		index++
		if index == len(gas) {
			index = 0
		}
		currentGas = currentGas - currentCost
		if currentGas < 0 {
			return false
		}
		if index == start && currentGas >= 0 {
			return true
		}
		currentGas += gas[index]
	}
}
