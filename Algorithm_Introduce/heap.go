package Algorithm_Introduce

// 堆就是一颗完全二叉树，一维数组就可以存储一个堆
// heap表示堆，size表示当前堆的大小，下标从1开始比较方便，l = root * 2, r = root * 2 + 1

// 堆的基本操作
// 1. 插入一个数，heap[++size] = k，up(size)
// 2. 求集合中的最小值，heap[1]，假设为小根堆
// 3. 删除最小值，heap[1] = heap[size]，size--，down(1)
// 4. 修改任意一个元素，假设要删除第k个点，heap[k] = heap[size]，size--，down(k)，up(k)
// 5. 修改任意一个元素，heap[k] = x, down(k), up(k)
// down和up同时只会有一个执行

var heap []int
var size int

// 以小根堆为例
func sort(nums []int) []int {
	n := len(nums)
	heap = make([]int, n+1)
	size = n

	for i := 1; i <= n; i++ {
		heap[i] = nums[i-1]
	}

	// 构造小根堆
	for i := n / 2; i > 0; i-- {
		down(i)
	}

	ans := make([]int, n)
	idx := 0

	for i := 1; i <= n; i++ {
		ans[idx] = heap[1]
		idx++

		heap[1] = heap[size]
		size--
		down(1)
	}
	return ans
}

// 下沉，当父节点的值 大于 孩子节点，将下沉
func down(u int) {
	t := u
	if u*2 <= size && heap[u*2] < heap[t] {
		t = u * 2
	}
	if u*2+1 <= size && heap[u*2+1] < heap[t] {
		t = u*2 + 1
	}
	if u != t {
		heap[u], heap[t] = heap[t], heap[u]
		down(t)
	}
}

// 上浮，孩子节点值 小于 父节点值，将上浮
func up(u int) {
	for u/2 > 0 && heap[u/2] > heap[u] {
		heap[u/2], heap[u] = heap[u], heap[u/2]
		u /= 2
	}
}
