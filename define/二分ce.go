package define

import (
	"fmt"
)

// 两数之和
func main() {
	n := 0
	fmt.Scanf("%d", &n)
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		temp := 0
		fmt.Scanf("%d", &temp)
		nums = append(nums, temp)
	}
	target := 0
	fmt.Scanf("%d", &target)

	ans := search(nums, target)
	fmt.Printf("%d", ans)
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { //[left,right]
		mid := (left + right) / 2
		if target > nums[mid] { //[left,mid,target,right]
			left = mid + 1
		} else if target < nums[mid] { //[left,target,mid,right]
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
