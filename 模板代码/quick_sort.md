### 代码    
1. 快速排序，使用中间值为pivot，使用端点值时，可能会超时，如果数据加强的话

```go
package main

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	j := partition(nums, left, right)
	quickSort(nums, left, j)
	quickSort(nums, j+1, right)
}

func partition(nums []int, left, right int) int {
	x := nums[left+(right-left)/2]
	i, j := left-1, right+1
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return j
}
```
