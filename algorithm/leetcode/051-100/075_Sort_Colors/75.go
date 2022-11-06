package _75_Sort_Colors

func sortColors(nums []int) {
	i, p0, p2 := 0, 0, len(nums)-1
	for i <= p2 {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
			continue
		} else if nums[i] == 1 {
			i++
			continue
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			continue
		}
	}
}
