package _54_4Sum_II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	calMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			calMap[sum]++
		}
	}
	res := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum := (nums3[i] + nums4[j]) * -1
			if value, ok := calMap[sum]; ok {
				res += value
			}
		}
	}
	return res
}
