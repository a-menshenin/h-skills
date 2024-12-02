package main

import "fmt"

func main(){
	res := longestConsecutive([]int{100,4,200,1,3,2})
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	m1 := make(map[int]struct{})
	for _, v := range nums {
		m1[v] = struct{}{}
	}



	var res int
	m := make(map[int]int)
	for k, _ := range m1 {
		if _, ok := m1[k-1]; ok {
			_, ok := m[k-1]
			_, ok2 := m[k+1]
			if ok || ok2 {
				m[k] = m[k-1] + 1
				if m[k] > res {
					res = m[k]
				}
			}

			m[k] = 1
		} else {
			m[k] = 1
		}
	}

	return res
}
