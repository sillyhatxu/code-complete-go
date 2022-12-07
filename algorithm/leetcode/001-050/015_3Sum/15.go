package _15_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			//-4,[-1,-1],0,1,2    -1在i=1处已经判断，不需要重复判断
			continue
		}
		mid := i + 1
		right := len(nums) - 1
		for mid < right {
			sum := nums[i] + nums[mid] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				mid++
			} else {
				// val == 0
				res = append(res, []int{nums[i], nums[mid], nums[right]})
				for mid < right && nums[mid] == nums[mid+1] {
					//   [x x x 1 1 1 x x x x]
					//去重当前坐标↑
					mid++
				}
				mid++
				for mid < right && nums[right] == nums[right-1] {
					//   [x x x x x x 4 4 4 x]
					//          去重当前坐标↑
					right--
				}
				right--
			}
		}
	}
	return res
}
