package main

import (
	"fmt"
	"slices"
)

func main(){
	moveZeroes2([]int{0,1,0,3,12})
}

func moveZeroes2(nums []int) {
	res := make([]int, 0, len(nums))
	j := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if j < 0 {
			break
		}

		// по i ищем только 0
		if nums[i] == 0 {
			if len(res) != 0 {
				res = append([]int{0}, res...)
			} else {
				res = []int{0}
			}
		}

		// По j ищем только не 0
		if nums[j] != 0 {
			res = append(res, nums[j])
		}

		j--
	}

	slices.Reverse(res)

	nums = res

	fmt.Println(nums)
}
