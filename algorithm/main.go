package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	count := RemoveDuplicates(nums)
	fmt.Println("count:", count)
}

func RemoveDuplicates(nums []int) int {
	numLen := len(nums)
	for i := 0; i < numLen; i++ {
		for k := i + 1; k < numLen; k++ {
			// 如果有相同的就从数组里删掉
			if nums[i] == nums[k] {
				nums = arrRemove(nums, i)
				// 删掉后数组长度-1
				numLen -= 1
			}
		}
	}
	fmt.Println("nums:", nums)
	return len(nums)
}

// 删除数组里第i个元素
func arrRemove(nums []int, i int) []int {
	if len(nums) >= i+1 {
		nums = append(nums[:i], nums[i+1:]...)
		return nums
	}
	return nums
}
