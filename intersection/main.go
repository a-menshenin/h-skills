package main

import (
	"fmt"
)

func main() {
	res := intersection2(
		[]int{1,2,2,1},
		[]int{2,2},
	)
	fmt.Println(res)
	res = intersection2(
		[]int{4,9,5},
		[]int{9,4,9,8,4},
	)
	fmt.Println(res)
}

func intersection(nums1 []int, nums2 []int) []int {
	uniqueNums1 := make(map[int]struct{})
	for _, v := range nums1 {
		uniqueNums1[v] = struct{}{}
	}

	intersectionMap := make(map[int]struct{})
	intersection := []int{}

	for _, v := range nums2 {
		if _, ok := uniqueNums1[v]; ok {
			if _, ok := intersectionMap[v]; !ok {
				intersectionMap[v] = struct{}{}
				intersection = append(intersection, v)
			}
		}
	}

	return intersection
}

func intersection2(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}

	set2 := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := set1[v]; ok {
			set2[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(set2))
	for v := range set2 {
		result = append(result, v)
	}

	return result
}
