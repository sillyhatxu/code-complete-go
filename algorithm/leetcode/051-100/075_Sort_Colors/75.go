package _75_Sort_Colors

func sortColors(nums []int) {
	i, color0, color2 := 0, 0, len(nums)-1
	for i <= color2 {
		if nums[i] == 0 {
			nums[i], nums[color0] = nums[color0], nums[i]
			color0++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[color2] = nums[color2], nums[i]
			color2--
		}
	}
}

func sortColors1(nums []int) {
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
