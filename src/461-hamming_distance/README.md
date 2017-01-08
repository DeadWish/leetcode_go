# 461. Hamming Distance

## 结果

0ms, 击败显示14%（这里肯定不对）

## 解法

一般跟bit相关的算法，都要想到二进制操作，与、或、异或等。此题是求异或后位为1的数量，那一定要先求异或，然后就此题变为：求一个二进制数1的位数。

1. 暴力解法1

   + 循环判断每位是否为1。


   + time:O(n), n为位数

2. 暴力解法2

   ```go
   input: a; output: i
   for a != 0 {
     a &= a - 1
     i++
   }
   return i
   ```

   + 我们发现`a &= a - 1`可以将最后为1的位给置零，切只用了一次与运算，所以我们可以用此方法只计算为1的位。
   + time:O(k), k为位为1的数量

3. 二分法（平行法）

   + 思路就是按两位处理，最后把总和加一起。

     a. 例如8位的，11111111，先按1位到2位处理，1+1, 1+1, 1+1, 1+1，结果就为10 10 10 10

     b. 再按2位到4位处理，2+2, 2+2，结果为0100 0100

     c. 再按4位到8位处理，4+4，结果为00001000

     d. 转换成10进制，8

   + 代码如下（32位）:

   ```go
   input, output: u
   u = (u & 0x55555555) + ((u >> 1) & 0x55555555);
   u = (u & 0x33333333) + ((u >> 2) & 0x33333333);
   u = (u & 0x0F0F0F0F) + ((u >> 4) & 0x0F0F0F0F);
   u = (u & 0x00FF00FF) + ((u >> 8) & 0x00FF00FF);
   u = (u & 0x0000FFFF) + ((u >> 16) & 0x0000FFFF);
   return u
   ```

   + 此思路一定要按照二分的方式理解
   + time:O(logn), n为位数

4. 查表法

   + 空间换效率
   + 例如8位的，就做一个256长度的表，32位的就做一个65535的表，更长的效率可能就没有计算高了。这里都做静态表，因为表是已知（可预先计算）
   + 也可以分段做查表计算，例如32位的使用256长度的表，分成4个8位的段，分别查表
   + time:O(k), space:O(j), k为分段数，j为表长 