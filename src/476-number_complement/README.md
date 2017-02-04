# 476. Number Complement

## 结果

0ms, 击败24.44%

## 解法

1. 按位求反，取相同有效位数即可。
	- 如果某个数非常大，例如二进制表示中1在31位，这样有效位数是31位，查找有效位数就要循环31次
	- time:O(n), space:O(1), n为int的位数

2. 可以先找到有效位数，然后使用有效位数个1和原来的数做异或操作。
	- 参见解题代码, [相关解释](https://discuss.leetcode.com/topic/74897/maybe-fewest-operations/6)
	- time:O(1), space:O(1),