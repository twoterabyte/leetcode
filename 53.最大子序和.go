/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// dp[i] 表示以 nums[i] 为结尾的连续子数组最大和
// dp[0] = nums[i]
// dp[n] = if dp[n-1]<0 then nums[i]
//         else dp[n-1] + nums[i]
// @lc code=start
func maxSubArray(nums []int) int {
	// 状态压缩，因为只依赖 dp 数组上一个元素
	pre := nums[0]
	result := pre
	for i := 1; i < len(nums); i++ {
		if pre < 0 {
			pre = nums[i]
		} else {
			pre = pre + nums[i]
		}
		if pre > result {
			result = pre
		}
	}
	return result
}
// @lc code=end

