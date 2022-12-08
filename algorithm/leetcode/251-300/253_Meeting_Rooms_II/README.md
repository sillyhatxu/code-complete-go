# 253. Meeting Rooms II

Question Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
find the minimum number of conference rooms required.

For example, Given [[0, 30],[5, 10],[15, 20]], return 2.

### Translate:

> 253. 会议室 II

给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。

### 示例 1：

```
输入：intervals = [[0,30],[5,10],[15,20]]
输出：2
```

### 示例 2：

```
输入：intervals = [[7,10],[2,4]]
输出：1
```

### 提示：

* 1 <= intervals.length <= 10^4 
* 0 <= start[i] < end[i] <= 10^6


### 解题思路
本题可以先假设只有一个大大大会议室，参加不同会议的人都可以自由进出会议室，咱们只要看到同一时间会议室内最多有几批人，就知道需要几个会议了。

咱们的做法是：

1、先把进出时间从小到大进行排序。
这里有个小技巧，怎么样区分哪个是进，哪个是出呢？只要加个标志位就好了。
为了不影响原数大小，所有数值*10后，再加上标志位，进为2，出为1。
为什么进要大一点呢？因为如果A组人员刚出，B组就进来了，那么可以认为一个会议室是可以容纳这两组人开会的。所以，咱们先出后进，防止多算。

2、如果有人进房间，那么需要的会议室数量curNeedRoom就+1，反之，有人出房间，curNeedRoom就-1。

3、记录下curNeedRoom的最大值，就是咱们需要的会议室数量
