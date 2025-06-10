from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        n = len(nums)

        nums.sort()

        return nums[n//2]

    def majorityElement_v2(self, nums: List[int]) -> int:
        """
        Given an array nums of size n, return the majority element.

        The majority element is the element that appears more than ⌊n / 2⌋ times. 
        You may assume that the majority element always exists in the array.

        Constraints:
            - n == nums.length
            - 1 <= n <= 5 * 10^4
            - -10^9 <= nums[i] <= 10^9

        """
        count = 0
        candidate = None
        for num in nums:
            if count == 0:
                candidate = num
            if num == candidate:
                count += 1
            else:
                count -= 1
        return candidate


if __name__ == "__main__":
    # Example usage
    solution = Solution()

    a1 = solution.majorityElement([3, 2, 3])
    assert 3 == a1

    a2 = solution.majorityElement([2, 2, 1, 1, 1, 2, 2])
    assert 2 == a2

    print(a1, a2)
