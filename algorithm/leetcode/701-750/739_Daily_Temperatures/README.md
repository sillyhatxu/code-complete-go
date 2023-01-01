# 739. Daily Temperatures

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

### Example 1:

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

### Example 2:

```
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
```

### Example 3:

```
Input: temperatures = [30,60,90]
Output: [1,1,0]
```

### Constraints:

* 1 <= temperatures.length <= 105
* 30 <= temperatures[i] <= 100

### Translate:

> 739. 每日温度

给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。

### 示例 1:

```
输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]
```

### 示例 2:

```
输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]
```

### 示例 3:

```
输入: temperatures = [30,60,90]
输出: [1,1,0]
```

### 提示：

* 1 <=temperatures.length <= 105
* 30 <=temperatures[i]<= 100


# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

|||||||
|:---:|:---:|:---:|:---:|:---:|:---:|
| i | 0 | 1 | 2 | 3 | 4 |
| temperatures | 40 | 30 | 50 | 20 | 30 |
| res | 2 | 1 | 0 | 1 | 0 |
```

# Complexity
- Time complexity: O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: O(n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func dailyTemperatures(temperatures []int) []int {
	var queue []int
	res := make([]int, len(temperatures), len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(queue) != 0 && temperatures[queue[len(queue)-1]] < temperatures[i] {
			idx := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			res[idx] = i - idx
		}
		queue = append(queue, i)
	}
	return res
}
```