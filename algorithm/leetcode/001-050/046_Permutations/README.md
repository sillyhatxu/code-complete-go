# 46. Permutations

Given an array `nums` of distinct integers, return all the possible permutations. You can return the answer in **any order**.


### Example 1:

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

### Example 2:

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

### Example 3:

```
Input: nums = [1]
Output: [[1]]
```

### Constraints:

* 1 <= nums.length <= 6
* -10 <= nums[i] <= 10
* All the integers of nums are unique.

### Translate:

> 46. 全排列
> 
> 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
> 
> 提示
> 
> 1 <= nums.length <= 6
> 
> -10 <= nums[i] <= 10
> 
> nums 中的所有整数 互不相同
>

### 解题思路

1. 先画一个树
```
                []
        /       |         \\
      1         2         3        
    /  |      /  |      /  |
   12  13   21   23   31   32
    |   |    |   |     |    |
   123  132  213 231  312  321

```

2. 深度优先遍历
   result 返回的数组
   currentNums当前未选中的数字数组
   candidates 已经选中的数字数组

* 第一步，result还是空的，先拿出1遍历，所以candidates = [1]; currentNums = [2,3]
* 第二步，result还是空的，candidates = [1，2]; currentNums = [3]
* 第三步，candidates = [1，2，3]; currentNums = []，因为已经没有数字需要选择了，所以把选好的数字放入result
* 第四步，回到上一步，这时result = [[1,2,3]];candidates = [1，3]; currentNums = [2]，重新迭代

### 代码

```golang
func permute(nums []int) [][]int {
	var result [][]int
	dfs(&result, nums, []int{})
	return result
}

func dfs(result *[][]int, currentNums []int, candidates []int) {
	if len(currentNums) == 0 {
		*result = append(*result, candidates)
		return
	}
	for i := 0; i < len(currentNums); i++ {
		temp := currentNums[i]
		dfs(result,append(append([]int{}, currentNums[0:i]...), currentNums[i+1:]...),append(candidates, temp))
	}
}
```



