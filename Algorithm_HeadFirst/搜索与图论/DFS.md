## DFS

DFS，又叫深度优先遍历、回溯，俗称暴力搜索

最重要的是需要思考用什么顺序来遍历



回溯，一定要注意恢复现场

往回走的过程叫回溯

```go
func dfs(xxx) {
    // for循环是横向的，for循环里面的dfs是纵向的
    for i := x; i < n; i++ {
        path = append(path, xx)
        dfs(xxx)
        path = path[:len(path)-1]
    }
}
```



## 回溯能解决什么问题

1. 组合问题
2. 切割问题
3. 子集问题
4. 排列问题
5. 棋盘问题
   1. 解数独
   2. N皇后







## 总结

1. 回溯，需要注意的是for循环中i从多少开始，递归的时候index传的是啥。注意：for循环不是必须得，但是大部分情况下都有for循环
2.  backtrace(nums, i+1, path)和backtrace(nums, index+1, path)，这两者的区别，好好体会
   1. 提示：index+1，不管i走到哪里，比如i走到index+3的位置，那下一次循环，i还是会从index+1开始
3.  []int 添加到 [][]int，应该怎么处理（golang）
4. 在dfs方法中，如果有for循环，那么i是从0开始，还是从index开始；下一次递归时入参是i+1，还是index+1，或者是其他？
5. 回溯三部曲
   1. 递归的参数和返回值
      1. 参数中的startIndex，其实就是表示下一层从哪里取数
   2. 确定终止条件
   3. 单层搜索的逻辑