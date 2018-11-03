package list

import "fmt"

func EmptyAppend() {
	var nums []int

	nums = append(nums, 1.0)
	nums = append(nums, 2.0)
	fmt.Println(nums)
}
