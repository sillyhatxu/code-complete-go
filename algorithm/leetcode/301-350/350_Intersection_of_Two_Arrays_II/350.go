package _50_Intersection_of_Two_Arrays_II

//Time: O(m+n); Space: O(min(m,n))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		intersect(nums2, nums1)
	}
	numMap := make(map[int]int)
	for _, num := range nums1 {
		numMap[num]++
	}
	var res []int
	for _, num := range nums2 {
		if value, ok := numMap[num]; ok && value > 0 {
			numMap[num]--
			res = append(res, num)
		}
	}
	return res
}
