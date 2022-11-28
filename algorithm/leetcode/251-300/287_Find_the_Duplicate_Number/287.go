package _87_Find_the_Duplicate_Number

func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if slow == fast {
			break
		}
	}
	check := 0
	for {
		slow = nums[slow]
		check = nums[check]
		if slow == check {
			break
		}
	}
	return slow
}
