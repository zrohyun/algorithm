from typing import List


class Solution:
    def containsNearbyDuplicate_v2(self, nums: List[int], k: int) -> bool:
        l_nums = len(nums)
        for i in range(l_nums):
            for j in range(i+1, l_nums):
                if j > i + k:
                    break
                if nums[i] == nums[j]:
                    return True

        return False

    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        """
        Given an integer array nums and an integer k,
        return true if there are two distinct indices i and j in the array
        such that nums[i] == nums[j] and abs(i - j) <= k.

        Example 1:

        Input: nums = [1,2,3,1], k = 3
        Output: true
        Example 2:

        Input: nums = [1,0,1,1], k = 1
        Output: true
        Example 3:

        Input: nums = [1,2,3,1,2,3], k = 2
        Output: false
        Constraints:
        - 1 <= nums.length <= 10^5
        - -10^9 <= nums[i] <= 10^9
        - 0 <= k <= 10^5
        """
        numsMap = {}
        for idx, n in enumerate(nums):
            if (prev_idx := numsMap.get(n)) is not None and idx - prev_idx <= k:
                return True
            numsMap[n] = idx

        return False


if __name__ == "__main__":
    # Example usage
    solution = Solution()

    # Test case 1
    param = [1, 2, 3, 1]
    k = 3
    expected_result = True
    result = solution.containsNearbyDuplicate(param, k)
    assert expected_result == result, f"{expected_result}, {result}"

    # Test case 2
    param = [1, 0, 1, 1]
    k = 1
    expected_result = True
    result = solution.containsNearbyDuplicate(param, k)
    assert expected_result == result, f"{expected_result}, {result}"

    # Test case 3
    param = [1, 2, 3, 1, 2, 3]
    k = 2
    expected_result = False
    result = solution.containsNearbyDuplicate(param, k)
    assert expected_result == result, f"{expected_result}, {result}"

    print("All test cases passed!")
