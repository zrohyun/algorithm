package main

import (
	"fmt"
)

// You are given two positive integers n and limit.
//
// Return the total number of ways to distribute n candies
// among 3 children such that no child gets more than limit candies.
//
// Constraints:
// 	- 1 <= n <= 10^6
// 	- 1 <= limit <= 10^6
func distributeCandies(n int, limit int) int {
	ans := 0
	for i := 0; i <= min(limit, n); i++ {
		if n-i > 2*limit {
			// 나머지 두 아이가 모든 사탕을 받을 수 없는 경우
			continue
		}
		ans += min(n-i, limit) - max(0, n-i-limit) + 1
	}
	return ans
}

func main() {
	a1 := distributeCandies(5, 2)
	if a1 != 3 {
		panic(fmt.Sprintf("Expected 3, got %d", a1))
	}

	a2 := distributeCandies(3, 3)
	if a2 != 10 {
		panic(fmt.Sprintf("Expected 10, got %d", a2))
	}
}
