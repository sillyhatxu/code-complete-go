# 1200. Minimum Absolute Difference

Given an array of **distinct** integers `arr`, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair `[a, b]` follows

* `a`, `b` are from `arr`
* `a` < `b`
* `b` - `a` equals to the minimum absolute difference of any two elements in `arr`



### Example 1:

```
Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
```

### Example 2:

```
Input: arr = [1,3,6,10,15]
Output: [[1,3]]
```

### Example 3:

```
Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]
```

### Constraints:

* 2 <= arr.length <= 105
* -106 <= arr[i] <= 106

### Translate:

> 1200. 最小绝对差
> 
> 给你个整数数组arr，其中每个元素都 不相同。
>
> 请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。
>
> 每对元素对 [a,b] 如下：
>
> a ,b均为数组arr中的元素
> 
> a < b
> 
> b - a等于 arr 中任意两个元素的最小绝对差