package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	name string
	prepare func() []int
	expected []int
}

func TestLRUCache(t *testing.T) {
	testCases := []testCase {
		// {
		// 	name: "case 1",
		// 	expected: []int{-1},
		// 	prepare: func() []int {
		// 		lru := NewLRUCache(2)
		// 		lru.Put(2, 1)
		// 		lru.Put(1, 1)
		// 		lru.Put(2, 3)
		// 		lru.Put(4, 1)

		// 		return []int{lru.Get(1)}
		// 	},
		// },
		// {
		// 	name: "case 2",
		// 	expected: []int{2,1,1,-1,3},
		// 	prepare: func() []int {
		// 		lru = NewLRUCache(2)
		// 		lru.data = map[int]int{}
		// 		lru.Put(2, 1)
		// 		lru.Put(3, 2)
		// 		v1 := lru.Get(3)
		// 		v2 := lru.Get(2)
		// 		lru.Put(4, 3)
		// 		v3 := lru.Get(2)
		// 		v4 := lru.Get(3)
		// 		v5 := lru.Get(4)

		// 		return []int{v1, v2, v3, v4, v5}
		// 	},
		// },
		{
			name: "case 3",
			expected: []int{-1,19,17,-1,-1,-1,5,-1,12,3,5,5,1,-1,30,5,30,-1,-1,24,18,-1,18,-1,18,-1,4,29,30,12,-1,29,17,22,18,-1,20,-1,18,18,20},
			prepare: func() []int {
				commands := []string{"LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"}
				intOutput := parseOutput("[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]")
				_, res := setLRUCacheAndGetResults(commands, intOutput)

				return res
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.prepare()
			for i, expectedValue := range tt.expected {
				if res[i] != expectedValue {
					t.Errorf("\nexpected res: %v \nactual res: %v", tt.expected, res)
				}
			}
		})
	}
}

func parseOutput(output string) [][]int {
	// Example: output = [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	pairs := strings.Split(strings.Trim(output, "[]"), "]")

	for i, pair := range pairs {
		pairs[i] = strings.Trim(pair, ",[")
	}

	res := make([][]int, 0, len(pairs))
	for _, pair := range pairs {
		strParts := strings.Split(pair, ",")

		intParts := make([]int, 0, len(strParts))
		for _, v := range strParts {
			i, _ := strconv.Atoi(v)
			intParts = append(intParts, i)
		}

		res = append(res, intParts)
	}

	return res
}

func setLRUCacheAndGetResults(commands []string, data [][]int) (*LRUCache, []int) {
	var (
		lru *LRUCache
		res []int
	)
	
	for i, cmd := range commands {
		switch cmd {
		case "LRUCache":
			lru = NewLRUCache(data[i][0])
		case "put":
			if data[i][0] == 11 && data[i][1] == 4 {
				fmt.Println("!!!")
			}
			lru.Put(data[i][0], data[i][1])
		case "get":
			res = append(res, lru.Get(data[i][0]))
		}
	}

	return lru, res
}
