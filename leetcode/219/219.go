package main

import (
	"fmt"
)

//Given an integer array nums and an integer k,
// return true if there are two distinct indices i and j in the array
// such that nums[i] == nums[j] and abs(i - j) <= k.
//
// Example 1:
//
// Input: nums = [1,2,3,1], k = 3
// Output: true
// Example 2:
//
// Input: nums = [1,0,1,1], k = 1
// Output: true
// Example 3:
//
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false
// Constraints:
// - 1 <= nums.length <= 10^5
// - -10^9 <= nums[i] <= 10^9
// - 0 <= k <= 10^5
func containsNearbyDuplicate_v1(nums []int, k int) bool {
    numsMap := map[int]int{}

    for idx, v := range nums {
        if _, ok := numsMap[v]; !ok{
            numsMap[v] = idx
        } else {
            if idx - numsMap[v] <= k {
                return true
            }else {
                numsMap[v] = idx
            }
        }
    }
    return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
    numsMap := map[int]int{}

    for idx, v := range nums {
        if prevIdx, exists := numsMap[v]; exists && idx - prevIdx <= k{
			return true
        } else {
			numsMap[v] = idx
        }
    }
    return false
}

func main() {
	// Test case 1
	param := []int{1,2,3,1}
	k := 3
	result := containsNearbyDuplicate(param,k)
	expected_result := true
	if result != expected_result {
	    fmt.Printf("Test case 1 failed: Expected %v, got %v\n", expected_result, result)
	}

	// Test case 2
	param = []int{1,0,1,1}
	k = 2
	result = containsNearbyDuplicate(param,k)
	expected_result = true
	if result != expected_result {
	    fmt.Printf("Test case 2 failed: Expected %v, got %v\n", expected_result, result)
	}

	// Test case 3
	param = []int{1,2,3,1,2,3}
	k = 2
	result = containsNearbyDuplicate(param,k)
	expected_result = false
	if result != expected_result {
	    fmt.Printf("Test case 3 failed: Expected %v, got %v\n", expected_result, result)
	}

	fmt.Println("All test cases passed!")
}
