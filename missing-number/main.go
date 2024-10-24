package main

import "fmt"

func main() {
	res := missingNumber([]int{3,0,1})

	fmt.Println(res)
}

func missingNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return -1
}
