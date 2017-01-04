# Two Sum

## 结果

9ms, 击败69.5%

## 解法

因为有Editorial Solution, 就简单总结一下

1. Brute Force，暴力破解。
   - 双循环
   - time:O(n^2), space:O(1)
2. Two-pass Hash Table, 两次哈希
   - 一次哈希记录所有输入，一次查找哈希
   - time:O(n), space:O(n)
3. One-pass Hash Table, 一次哈希
   - 在解法2中，记录的过程就可以检查，所以不一定要哈希所有输入，所以会比解法2更有效率
   - time:O(n), space:O(n)