### 代码   

```go
package main

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	tmp := make([]int, r-l+1)
	k := 0
	i, j := l, mid+1
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = nums[j]
		k++
		j++
	}
	for i = 0; i < k; i++ {
		nums[l+i] = tmp[i]
	}
}
```