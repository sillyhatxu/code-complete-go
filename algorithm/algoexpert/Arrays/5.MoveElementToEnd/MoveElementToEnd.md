# Move Element To End

**Difficulty:**Medium; **Category:**Arrays

You're given an array of integers and an integer. Write a function that moves
all instances of that integer in the array to the end of the array and returns
the array.

The function should perform this in place (i.e., it should mutate the input
array) and doesn't need to maintain the order of the other integers.

### Sample Input

```
array = [2, 1, 2, 2, 2, 3, 4, 2]
toMove = 2
```

### Sample Output

```
[1, 3, 4, 2, 2, 2, 2, 2] // the numbers 1, 3, and 4 could be ordered differently
```

### Hints


#### Hint 1

```
You can solve this problem in linear time.
```

#### Hint 2

```
In view of Hint #1, you can solve this problem without sorting the input array. Try setting two pointers at the start and end of the array, respectively, and progressively moving them inwards.
```

#### Hint 3

```
Following Hint #2, set two pointers at the start and end of the array, respectively. Move the right pointer inwards so long as it points to the integer to move, and move the left pointer inwards so long as it doesn't point to the integer to move. When both pointers aren't moving, swap their values in place. Repeat this process until the pointers pass each other.
```

#### Optimal Space & Time Complexity

```
O(n) time | O(1) space - where n is the length of the array
```

### Hints


## Test


### Test Case 1

```
{"array": [2, 1, 2, 2, 2, 3, 4, 2], "toMove": 2}
```

### Test Case 2

```
{"array": [], "toMove": 3}
```

### Test Case 3

```
{"array": [1, 2, 4, 5, 6], "toMove": 3}
```

### Test Case 4

```
{"array": [3, 3, 3, 3, 3], "toMove": 3}
```

### Test Case 5

```
{"array": [3, 1, 2, 4, 5], "toMove": 3}
```

### Test Case 6

```
{"array": [1, 2, 4, 5, 3], "toMove": 3}
```

### Test Case 7

```
{"array": [1, 2, 3, 4, 5], "toMove": 3}
```

### Test Case 8

```
{"array": [5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12], "toMove": 5}
```

### Test Case 9

```
{"array": [1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5], "toMove": 5}
```

### Test Case 10

```
{"array": [5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12], "toMove": 5}
```
