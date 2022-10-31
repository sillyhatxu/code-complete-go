# 31. Next Permutation

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

```
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

### Translate:

> 给定一个数组 [1,2,3] 这个数组可以组成一个数字 123，把可以组成的所有的数字，按照从小到大排列，123这个数字的下一个数字就是132


### 解题思路

1. len == 0 || len == 1 直接return
2. 如果本身是单调递减(321)直接反转即可(123)
3. 正式判断
   
3.1 先从右往定位坐标，找到第一个破坏单调递增的坐标x，如数字：12395521 -> [1, 2, 3, 9, 5, 5, 2, 1]
```
  x    
123 | 95521
```

3.2 再从右往定位坐标，找到第一个比x大的值坐标y
```
  x    
123 | 95 5 21
         y
```

3.3 交换x和y的值
```
  x    
125 | 95 3 21
         y
```

3.4 把x后边的值反转，最后得出12512359
```
  x    
125 | 12359
```

### 代码

```golang
func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	x := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			x = i
			break
		}
	}
	if x == -1 {
		reverse(&nums, 0, len(nums)-1)
		return
	}
	y := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[x] {
			y = i
			break
		}
	}
	nums[x], nums[y] = nums[y], nums[x]
	reverse(&nums, x+1, len(nums)-1)
}

func reverse(nums *[]int, start int, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}
```