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

	for {
		nums := make([]int, 10000)
		defer func() {
			nums = nums[1:]
		}()
	}
}
