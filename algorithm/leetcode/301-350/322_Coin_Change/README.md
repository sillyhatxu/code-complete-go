# 322. Coin Change

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

### Example 1:

```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
```

### Example 2:

```
Input: coins = [2], amount = 3
Output: -1
```

### Example 3:

```
Input: coins = [1], amount = 0
Output: 0
```

### Constraints:

* 1 <= coins.length <= 12
* 1 <= coins[i] <= 231 - 1
* 0 <= amount <= 104

### Translate:

> 322. 零钱兑换

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。


### 解题思路

#### eg1:
Input: coins = [1,2,5], amount = 6


* i : 我们需要获取的总金额（amount）
1. 初始化
```
dp[0] = 0
[ 0  7  7  7  7  7  7]
```
2. dp[1] 
```
When coin = 1; i(1)-1=0   dp[0]+1=2    dp[2] = 2
When coin = 2; i(1)-2=-1  无效
When coin = 5; i(1)-5=-4  无效

dp[1] = 1
[ 0  1  7  7  7  7  7]
```
3. dp[2]
```
When coin = 1; i(2)-1=1  dp[1] + 1 = 2    dp[2] = 2
When coin = 2; i(2)-2=0  dp[0] + 1 = 1    dp[2] = 1
When coin = 5; i(2)-5=-3 无效

dp[2] = min(1,2) = 1
[ 0  1  1  7  7  7  7]
```
4. dp[3]
```
                      ↓-----↓    ↓this is current coin
When coin = 1; i(3)-1=2  dp[2] + 1 = 3     dp[3] = 3
When coin = 2; i(3)-2=1  dp[1] + 1 = 2     dp[3] = 2
When coin = 5; i(3)-5=-2 无效

dp[3] = min(2,3) = 2
[ 0  1  1  2  7  7  7]
```

4. dp[4]
```
                      ↓-----↓    ↓this is current coin
When coin = 1; i(4)-1=3  dp[3] + 1 = 3   dp[4] = 3
When coin = 2; i(4)-2=2  dp[1] + 1 = 2   dp[4] = 2
When coin = 5; i(4)-5=-1 无效

dp[4] = min(2,3) = 2
[ 0  1  1  2  2  7  7]
```

4. dp[5]
```
                      ↓-----↓    ↓this is current coin
When coin = 1; i(5)-1=4  dp[4] + 1 = 3   dp[5] = 3
When coin = 2; i(5)-2=3  dp[3] + 1 = 3   dp[5] = 3
When coin = 5; i(5)-5=0  dp[0] + 1 = 1   dp[5] = 1

dp[5] = min(1,3,3) = 1
[ 0  1  1  2  2  1  7]
```

5. dp[6]
```
                      ↓-----↓    ↓this is current coin
When coin = 1; i(6)-1=5  dp[5] + 1 = 2   dp[6] = 2
When coin = 2; i(6)-2=4  dp[3] + 1 = 3   dp[6] = 3
When coin = 5; i(6)-5=1  dp[1] + 1 = 2   dp[6] = 2

dp[6] = min(2,3,2) = 2
[ 0  1  1  2  2  1  2]
```

### 代码

```golang
func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	total := amount+1
	dp := make([]int, total,total)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = total
		for _, coin := range coins {
			if (i-coin) >= 0 && dp[i-coin] != total {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == total {
		return -1
	}
	return dp[amount]
}
```

### 解题思路

```
dp[i]表示凑成金额i所需的最少硬币个数
遍历硬币数组coins，对于每个硬币coin，从金额coin开始递推，更新dp数组。如果凑成金额i需要硬币coin，则有dp[i] = min(dp[i], dp[i-coin]+1)。
```

### index

```
凑硬币，总数
```