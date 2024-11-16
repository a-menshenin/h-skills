package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func main() {
	fmt.Println(firstBadVersion(2))
}

func isBadVersion(version int) bool {
	if version == 2 {
		return true
	}

	return false
}

func firstBadVersion(n int) int {
	return detectBadVersionInHalfPart(1, n, -1)
}

func detectBadVersionInHalfPart(start, end int, earliestBadVersion int) int {
	middle := (start + end) / 2
	if !isBadVersion(middle) {
		if start == end {
			return earliestBadVersion
		}
		// Перебираем вторую половину
		return detectBadVersionInHalfPart(middle+1, end, earliestBadVersion)
	} else {
		if start == middle {
			return middle
		}

		// Перебираем первую половину
		return detectBadVersionInHalfPart(start, middle-1, middle)
	}
}
