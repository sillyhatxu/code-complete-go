package _88_Merge_Sorted_Array

//0 ms 100%; Time: O(m+n); Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 || (p1 >= 0 && nums1[p1] >= nums2[p2]) {
			nums1[i] = nums1[p1]
			p1--
		} else if p1 < 0 || (p2 >= 0 && nums1[p1] < nums2[p2]) {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}
