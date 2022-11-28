package _83_Move_Zeroes

func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p] = nums[i]
			p++
		}
	}
	for i := p; i < len(nums); i++ {
		nums[i] = 0
	}
}
