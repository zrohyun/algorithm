package main

import (
	"fmt"
)

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.

// Constraints:
//     - 1 <= nums.length <= 10^5
//     - -10^9 <= nums[i] <= 10^9
func containsDuplicate(nums []int) bool {
   
	// compare length set nums and list nums
	nums_map := make(map[int]bool)
	for _, num := range nums {
		if _, exists := nums_map[num]; exists {
			return true
		}
		nums_map[num] = true
	}
	return false
}

func main() {
	nums := []int{1,2,3,1}
	a1 := containsDuplicate(nums)
	if a1 != true {
		fmt.Printf("Expected true, got %t\n", a1)
	}

	nums = []int{1,2,3,4}
	a2 := containsDuplicate(nums)
	if a2 != false {
		fmt.Printf("Expected true, got %t\n", a2)

	}	
	nums = []int{1,1,1,3,3,4,3,2,4,2}
	a3 := containsDuplicate(nums)
	if a3 != true {
		fmt.Printf("Expected true, got %t\n", a3)
	}	
}
