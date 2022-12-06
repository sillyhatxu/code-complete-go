package main

func maxArea(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left, right, maxContainer := 0, len(height)-1, 0
	for left < right {
		h := min(height[left], height[right])
		l := right - left
		maxContainer = max(maxContainer, h*l)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxContainer
}
