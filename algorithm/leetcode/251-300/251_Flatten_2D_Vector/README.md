# 251. Flatten 2D Vector

Implement an iterator to flatten a 2d vector.

For example, Given 2d vector =

```
[
  [1,2],
  [3],
  [4,5,6]
]
```
By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,2,3,4,5,6].

Follow up: As an added challenge, try to code it using only iterators in C++ or iterators in Java.

### Translate:

> 251. 展开二维向量

请设计并实现一个能够展开二维向量的迭代器。该迭代器需要支持next 和hasNext两种操作。

示例：

```
Vector2D iterator = new Vector2D([[1,2],[3],[4]]);

iterator.next(); // 返回 1
iterator.next(); // 返回 2
iterator.next(); // 返回 3
iterator.hasNext(); // 返回 true
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 4
iterator.hasNext(); // 返回 false
```

注意：

请记得重置在 Vector2D 中声明的类变量（静态变量），因为类变量会在多个测试用例中保持不变，影响判题准确。请 查阅 这里。
你可以假定 next() 的调用总是合法的，即当 next() 被调用时，二维向量总是存在至少一个后续元素。

进阶：尝试在代码中仅使用 C++ 提供的迭代器 或 Java 提供的迭代器。