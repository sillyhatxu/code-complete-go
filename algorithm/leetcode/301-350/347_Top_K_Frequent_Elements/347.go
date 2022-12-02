package _47_Top_K_Frequent_Elements

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		if _, ok := frequency[num]; ok {
			frequency[num]++
		} else {
			frequency[num] = 1
		}
	}
	bucket := make([][]int, len(nums)+1, len(nums)+1)
	for k, v := range frequency {
		bucket[v] = append(bucket[v], k)
	}
	var res []int
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		res = append(res, bucket[i]...)
		if len(res) == k {
			break
		}
	}
	return res
}
