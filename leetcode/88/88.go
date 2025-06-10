package main

import "sort"

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function,
// but instead be stored inside the array nums1.
//
// To accommodate this, nums1 has a length of m + n,
// where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
func merge(nums1 []int, m int, nums2 []int, n int) {
    for j, i := 0, m; j < n; j, i = j+1, i+1 {
        nums1[i] = nums2[j]
    }
    sort.Ints(nums1)
}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
    merge(nums1[:], m, nums2[:], n)

}
