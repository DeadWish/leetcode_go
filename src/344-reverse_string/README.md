# 344. Reverse String

## 结果

6ms, 72.1%

## 解法

Go语言没提供直接转换的方法，所以选择将string转换为slice，这里代码里投机取巧了，用了[]byte，在正常使用的时候推荐使用[]rune，要不然utf8字符无法处理