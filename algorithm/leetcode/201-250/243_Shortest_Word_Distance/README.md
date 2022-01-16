# 243. Shortest Word Distance

Given an array of strings wordsDict and two different strings that already exist in the array word1 and word2, return the shortest distance between these two words in the list.

### Example 1:

```
Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "coding", word2 = "practice"
Output: 3
```

### Example 2:

```
Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
Output: 1
```

### Constraints:

* 1 <= wordsDict.length <= 3 * 104
* 1 <= wordsDict[i].length <= 10
* wordsDict[i] consists of lowercase English letters.
* word1 and word2 are in wordsDict.
* word1 != word2

### Translate:

> 两个单词在传入数组中的最小距离（注意第二个例子，有两个makes，coding距离第一个makes的距离是2，距离第二个makes的距离是1，所以返回1）