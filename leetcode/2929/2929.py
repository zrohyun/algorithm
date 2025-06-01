class Solution:
    def distributeCandies(self, n: int, limit: int) -> int:
        """
        You are given two positive integers n and limit.

        Return the total number of ways to distribute n candies
        among 3 children such that no child gets more than limit candies.

        Constraints:
            - 1 <= n <= 10^6
            - 1 <= limit <= 10^6
        """
        ans = 0
        for i in range(min(limit, n) + 1):
            if n - i > 2 * limit:
                # 나머지 두 아이가 모든 사탕을 받을 수 없는 경우
                continue
            ans += min(n - i, limit) - max(0, n - i - limit) + 1
        return ans


if __name__ == "__main__":
    # Example usage
    solution = Solution()

    a1 = solution.distributeCandies(5, 2)
    assert 3 == a1

    a2 = solution.distributeCandies(3, 3)
    assert 10 == a2
    print(a1, a2)
