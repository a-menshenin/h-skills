package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := findNthDigit(13)

	fmt.Println(res)
}

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	//1,2,3,4,5,6,7,8,9,10,11,12,13,14

	
	for i, j := 1, 0; i <= n; {
		for _, v := range strconv.Itoa(i) {
			j++
			intVal, _ := strconv.Atoi(string(v))

			fmt.Println("intVal")
			fmt.Println(intVal)
			fmt.Println("i")
			fmt.Println(i)
			fmt.Println("j")
			fmt.Println(j)

			if j == n {
				return intVal
			}

			i++
		}
		fmt.Println("i after for")
		fmt.Println(i)
	}

	return -1
}
