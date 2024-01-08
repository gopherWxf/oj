package define

import (
	"fmt"
	"time"
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
	time.Sleep(10 * time.Second)
	fmt.Printf("%d", ans+1)
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
