package main

import (
	"fmt"
	"sort"
)

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//Constraints:
// n == nums.length
// 1 <= n <= 5 * 10^4
// -10^9 <= nums[i] <= 10^9
func majorityElement_v1(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate

}

func majorityElement(nums []int) int {
	nums_len := len(nums)
	sort.Ints(nums)
	return nums[int(nums_len/2)] 
}


func main() {
	a1 := majorityElement([]int{3, 2, 3})
	if a1 != 3 {
		fmt.Printf("Expected 3, got %d\n", a1)
	}

	a2 := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if a2 != 2 {
		fmt.Printf("Expected 2, got %d\n", a2)
	}
}
