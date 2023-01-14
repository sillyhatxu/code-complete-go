package _16_Partition_Equal_Subset_Sum

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1, target+1)
	dp[0] = true
	for _, num := range nums {
		for currSum := target; currSum > 0; currSum-- {
			//x + num = currSum 是否有一个值 x + num = Current Sum
			temp := currSum - num
			if temp < 0 {
				break
			} else if temp >= 0 {
				dp[currSum] = dp[currSum] || dp[temp]
			}
			if dp[target] {
				break
			}
		}
	}
	return dp[target]
}
