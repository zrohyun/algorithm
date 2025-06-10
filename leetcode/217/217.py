from typing import List


class Solution:

    def containsDuplicate(self, nums: List[int]) -> bool:
        """
        Given an integer array nums, return true if any value appears at least twice in the array, 
        and return false if every element is distinct.

        Constraints:
            - 1 <= nums.length <= 10^5
            - -10^9 <= nums[i] <= 10^9
        """

        return len(nums) != len(set(nums))


if __name__ == "__main__":
    # Example usage
    solution = Solution()

    nums = [1, 2, 3, 1]
    a1 = solution.containsDuplicate(nums)
    assert True == a1

    nums = [1, 2, 3, 4]
    a2 = solution.containsDuplicate(nums)
    assert False == a2

    nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
    a3 = solution.containsDuplicate(nums)
    assert True == a3

    print(a1, a2, a3)
