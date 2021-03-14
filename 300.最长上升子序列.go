/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// dp(i) 表示以第 i 个元素为结尾的递增子序列长度
// dp(0) = 1
// dp(n) = max{dp(i) + 1 for i in i < n && arr[i] < arr[n]} or 1
// @lc code=start
func lengthOfLIS(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}
	dp := make([]int, len(nums))
	max := 1
	for n, num := range nums {
		dp[n] = 1
		for i := 0; i < n; i++ {
			if nums[i] >= num {
				continue
			}
			if dp[i] + 1 > dp[n] {
				dp[n] = dp[i] + 1
			}
		}
		if dp[n] > max {
			max = dp[n]
		}
	}
	return max
}
// @lc code=end

